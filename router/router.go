package router

import (
	"ai-dekawaf/app/auditlog"
	"ai-dekawaf/app/f5"
	"ai-dekawaf/app/permission"
	"ai-dekawaf/app/project"
	"ai-dekawaf/app/securityprofile"
	"ai-dekawaf/app/user"
	"ai-dekawaf/database"
	"ai-dekawaf/middleware"
	"context"
	"fmt"
	"net/http"
	"os"

	_ "ai-dekawaf/docs"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
)

var (
	db = database.PostsqlConn()

	// repositories
	securityprofileRepo securityprofile.Repository = securityprofile.NewRepository(db)
	auditlogRepo        auditlog.Repository        = auditlog.NewRepository(db)
	f5Repo              f5.Repository              = f5.NewRepository()
	projectRepo         project.Repository         = project.NewRepository(db)
	userRepo            user.Repository            = user.NewRepository(db)

	// usecases
	securityprofileUsecase securityprofile.Usecase = securityprofile.NewUsecase(securityprofileRepo, auditlogRepo, f5Repo)
	permissionUsecase      permission.Usecase      = permission.NewUsecase(projectRepo, userRepo)

	// handler
	securityprofileHandler securityprofile.Handler = securityprofile.NewHandler(securityprofileUsecase, permissionUsecase, projectRepo)
)

func SetupRouter() (*gin.Engine, error) {
	if err := f5Repo.Login(context.Background()); err != nil {
		return nil, fmt.Errorf("initialize F5 authentication: %w", err)
	}

	// websocket.InitHandler()

	r := gin.Default()
	if os.Getenv("ENV") != "local" {
		gin.SetMode(gin.ReleaseMode)
	}

	r.Use(gin.Recovery())
	r.Use(CORSMiddleware())
	r.Use(middleware.PrometheusHTTPMetrics())
	r.GET("/metrics", middleware.PrometheusMetricsHandler())

	r.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/health-check", middleware.AuthorizeJWT(), func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Ok")
	})

	rSecurityProfile := r.Group("/security-profile")
	rSecurityProfile.GET("", securityprofileHandler.Page)
	rSecurityProfile.POST("", securityprofileHandler.Create)

	return r, nil
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", os.Getenv("ORIGIN"))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Permissions-Policy", "geolocation=(), camera=(), microphone=()")
		c.Writer.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		c.Writer.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
		c.Writer.Header().Set("Cross-Origin-Resource-Policy", "same-site")
		c.Writer.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
