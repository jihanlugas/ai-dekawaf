package securityprofile

import (
	"ai-dekawaf/request"
	"ai-dekawaf/response"
	"ai-dekawaf/validator"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leebenson/conform"
)

type Handler struct {
	Usecase Usecase
}

func NewHandler(usecase Usecase) Handler {
	return Handler{
		Usecase: usecase,
	}
}

// Page godoc
// @Summary Get security profiles
// @Description Get security profiles
// @Tags Security Profile
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param page query request.PageSecurityProfile false "Page number"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /security-profile [get]
func (h *Handler) Page(c *gin.Context) {
	var err error
	var req request.PageSecurityProfile
	err = c.ShouldBindQuery(&req)
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

	mSecurityProfiles, count, err := h.Usecase.PageSecurityProfile(c, req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to retrieve security profiles", err, nil)
		return
	}

	response.Success(c, http.StatusOK, "Security profile successfully retrieved", response.PayloadPagination(&req, mSecurityProfiles, count))
}

// Create godoc
// @Summary Create security profile
// @Description Create security profile
// @Tags Security Profile
// @Security ApiKeyAuth
// @Accept multipart/form-data
// @Produce json
// @Param projectId formData string true "Project ID"
// @Param name formData string true "Name"
// @Param description formData string false "Description"
// @Param domain formData string true "Domain"
// @Param node formData string true "Node"
// @Param pool formData string true "Pool"
// @Param port formData string true "Port"
// @Param destination formData string true "Destination (IPv4)"
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

// Get godoc
// @Summary Get security profile
// @Description Get security profile
// @Tags Security Profile
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "Security profile ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /security-profile/{id} [get]
func (h *Handler) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "Missing security profile ID", nil, nil)
		return
	}

	mSecurityProfile, err := h.Usecase.GetSecurityProfile(c, id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get security profile", err, nil)
		return
	}

	response.Success(c, http.StatusOK, "Security profile retrieved", mSecurityProfile)
}

// Delete godoc
// @Summary Delete security profile
// @Description Delete security profile
// @Tags Security Profile
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "Security profile ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /security-profile/{id}/delete [post]
func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "Missing security profile ID", nil, nil)
		return
	}

	err := h.Usecase.DeleteSecurityProfile(c, id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete security profile", err, nil)
		return
	}

	response.Success(c, http.StatusOK, "Security profile deleted", nil)
}
