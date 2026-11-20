package securityprofile

import (
	"ai-dekawaf/app/permission"
	"ai-dekawaf/app/project"
	"ai-dekawaf/request"
	"ai-dekawaf/response"
	"ai-dekawaf/validator"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leebenson/conform"
)

type Handler struct {
	Usecase           Usecase
	PermissionUsecase permission.Usecase
	ProjectRepo       project.Repository
}

func NewHandler(usecase Usecase, permissionUsecase permission.Usecase, projectRepo project.Repository) Handler {
	return Handler{
		Usecase:           usecase,
		PermissionUsecase: permissionUsecase,
		ProjectRepo:       projectRepo,
	}
}

// Page godoc
// @Summary Get security profiles
// @Description Get security profiles
// @Tags Security Profile
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /security-profile [get]
func (h *Handler) Page(c *gin.Context) {
	response.Success(c, http.StatusOK, "Security profile successfully retrieved", nil)
}

// Create godoc
// @Summary Create security profile
// @Description Create security profile
// @Tags Security Profile
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "Name"
// @Param description formData string false "Description"
// @Param domain formData string true "Domain"
// @Param node formData string true "Node"
// @Param pool formData string true "Pool"
// @Param port formData string true "Port"
// @Param serverTechnologies formData []string false "Server technologies" collectionFormat(multi)
// @Param enforcementMode formData string true "Enforcement mode"
// @Param botProtection formData bool false "Bot protection"
// @Param certificateFile formData file true "Certificate file"
// @Param keyFile formData file true "Key file"
// @Param chainFile formData file false "Certificate chain file"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /security-profile [post]
func (h *Handler) Create(c *gin.Context) {
	var err error
	var req request.CreateSecurityProfile

	err = c.ShouldBind(&req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request", err, nil)
		return
	}

	// Trim whitespaces.
	conform.Strings(&req)

	// Validate request body.
	v := validator.Validator
	err = v.Struct(&req)
	if err != nil {
		errValidator := validator.MapValidationErrors(err)
		response.Error(c, http.StatusBadRequest, "Invalid request", err, errValidator)
		return
	}

	err = h.Usecase.CreateSecurityProfile(c, req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create security profile", err, nil)
		return
	}

	response.Success(c, http.StatusOK, "Security profile created", nil)
}
