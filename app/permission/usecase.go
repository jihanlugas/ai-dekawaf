package permission

import (
	"dekawaf/app/project"
	"dekawaf/app/user"
	"dekawaf/constant"
	"dekawaf/middleware"
	"dekawaf/model"

	"github.com/gin-gonic/gin"
)

type Usecase interface {
	// ChecUserPermission checks user permission and checking user have access to the project.
	CheckUserPermission(c *gin.Context, projectID, privilegeName string, requireEditorAccess bool) bool

	// CheckUserPrivilege checks user privilege without checking user have access to the project.
	CheckUserPrivilege(c *gin.Context, privilegeName string, requireEditorAccess bool) bool
}

type permissionUsecase struct {
	projectRepo project.Repository
	userRepo    user.Repository
}

func NewPermissionUsecase(projectRepo project.Repository, userRepo user.Repository) Usecase {
	return &permissionUsecase{
		projectRepo: projectRepo,
		userRepo:    userRepo,
	}
}

// CheckUserPermission checks if the user has the required privileges to perform the action.
func (u *permissionUsecase) CheckUserPermission(c *gin.Context, projectID, privilegeName string, requireEditorAccess bool) bool {
	userID := middleware.GetUserID(c)
	userRole := middleware.GetUserRole(c)

	// Check privilege for user tenant.
	if userRole != constant.RoleSuperadmin && userRole != constant.RoleGodMode {
		// Check user have access to the project.
		projectUsers, err := u.projectRepo.GetProjectByUserIdAndProjectId(projectID, userID)
		if err != nil {
			return false
		}

		if projectUsers.UserID != userID {
			return false
		}

		// Get user by user ID.
		user, err := u.userRepo.GetUserByID(userID)
		if err != nil {
			return false
		}

		// When organization role is not set, the user have no privileges.
		if user.OrganizationRole == nil {
			return false
		}

		// Check if the user has the required privileges.
		return user.OrganizationRole.CheckForPrivilege(privilegeName, requireEditorAccess)
	}

	if userRole == constant.RoleSuperadmin {
		// Check if superadmin has the required privileges.
		superadminPrivileges, err := u.userRepo.GetSuperAdminPrivilegesByID(userID)
		if err != nil {
			return false
		}

		return model.CheckPrivilegeInRoles(superadminPrivileges, privilegeName, requireEditorAccess)
	}

	if userRole == constant.RoleGodMode {
		return true
	}

	return false
}

func (u *permissionUsecase) CheckUserPrivilege(c *gin.Context, privilegeName string, requireEditorAccess bool) bool {
	userID := middleware.GetUserID(c)
	userRole := middleware.GetUserRole(c)

	// Check privilege for user tenant.
	if userRole != constant.RoleSuperadmin && userRole != constant.RoleGodMode {
		// Get user by user ID.
		user, err := u.userRepo.GetUserByID(userID)
		if err != nil {
			return false
		}

		// When organization role is not set, the user have no privileges.
		if user.OrganizationRole == nil {
			return false
		}

		// Check if the user has the required privileges.
		return user.OrganizationRole.CheckForPrivilege(privilegeName, requireEditorAccess)
	}

	// Check if superadmin has the required privileges.
	superadminPrivileges, err := u.userRepo.GetSuperAdminPrivilegesByID(userID)
	if err != nil {
		return false
	}

	return model.CheckPrivilegeInRoles(superadminPrivileges, privilegeName, requireEditorAccess)
}
