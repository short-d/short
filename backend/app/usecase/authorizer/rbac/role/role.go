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
		permission.ViewShortLink,

		permission.ViewAdminPanel,
	},
	ShortLinkEditor: {
		permission.ViewShortLink,
		permission.CreateShortLink,
		permission.EditShortLink,
		permission.DisableShortLink,
		permission.DeleteShortLink,

		permission.ViewAdminPanel,
	},
	ChangeLogViewer: {
		permission.ViewChange,

		permission.ViewAdminPanel,
	},
	ChangeLogEditor: {
		permission.ViewChange,
		permission.CreateChange,
		permission.EditChange,
		permission.DeleteChange,

		permission.ViewAdminPanel,
	},
	SecuritySpecialist: {
		permission.DisableShortLink,
		permission.DisableUser,

		permission.ViewAdminPanel,
	},
	Admin: {
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

		permission.ViewAdminPanel,
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
