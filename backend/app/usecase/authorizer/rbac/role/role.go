package role

import "github.com/short-d/short/backend/app/usecase/authorizer/rbac/permission"

// Role contains a list of permissions
type Role string

const (
	Basic Role = "basic"

	SecuritySpecialist Role = "security_specialist"

	ShortLinkViewer Role = "short_link_viewer"
	ShortLinkEditor Role = "short_link_editor"

	ChangeLogViewer Role = "changelog_viewer"
	ChangeLogEditor Role = "changelog_editor"

	Admin Role = "admin"
)

var permissions = map[Role][]permission.Permission{
	Basic: {},
	ShortLinkViewer: {
		permission.ViewAdminPanel,

		permission.ViewShortLink,
	},
	ShortLinkEditor: {
		permission.ViewAdminPanel,

		permission.ViewShortLink,
		permission.CreateShortLink,
		permission.EditShortLink,
		permission.DisableShortLink,
		permission.DeleteShortLink,
	},
	ChangeLogViewer: {
		permission.ViewAdminPanel,

		permission.ViewChange,
	},
	ChangeLogEditor: {
		permission.ViewAdminPanel,

		permission.ViewChange,
		permission.CreateChange,
		permission.EditChange,
		permission.DeleteChange,
	},
	SecuritySpecialist: {
		permission.ViewAdminPanel,

		permission.DisableShortLink,
		permission.DisableUser,
	},
	Admin: {
		permission.ViewAdminPanel,

		permission.ViewShortLink,
		permission.CreateShortLink,
		permission.EditShortLink,
		permission.DisableShortLink,
		permission.DeleteShortLink,

		permission.ViewChange,
		permission.CreateChange,
		permission.EditChange,
		permission.DeleteChange,

		permission.UpgradeUser,
		permission.DowngradeUser,
		permission.DisableUser,
		permission.DeleteUser,

		permission.CreateAPIKey,
	},
}

// HasPermission checks whether a role grants the requested permission.
func (r Role) HasPermission(permission permission.Permission) bool {
	for _, value := range permissions[r] {
		if value == permission {
			return true
		}
	}
	return false
}
