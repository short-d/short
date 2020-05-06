package role

import "github.com/short-d/short/backend/app/usecase/authorizer/permission"

// Role represents a groupings of permissions
type Role string

const (
	Basic Role = "basic"

	SecuritySpecialist Role = "security_specialist"

	ShortLinkViewer Role = "short_link_viewer"
	ShortLinkEditor Role = "short_link_editor"

	ChangeLogViewer Role = "change_log_viewer"
	ChangeLogEditor Role = "change_log_editor"

	Admin Role = "admin"
)

var permissions = map[Role][]permission.Permission{
	Basic: {},
	ShortLinkViewer: {
		permission.ViewShortLink,
	},
	ShortLinkEditor: {
		permission.ViewShortLink,
		permission.CreateShortLink,
		permission.EditShortLink,
		permission.DisableShortLink,
		permission.DeleteShortLink,
	},
	ChangeLogViewer: {
		permission.ViewChange,
	},
	ChangeLogEditor: {
		permission.ViewChange,
		permission.CreateChange,
		permission.EditChange,
		permission.DeleteChange,
	},
	SecuritySpecialist: {
		permission.DisableShortLink,
		permission.DisableUser,
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

		permission.ViewDashboards,
	},
}

// HasPermission tells if the given role grants access to a permission
func (r Role) HasPermission(permission permission.Permission) bool {
	for _, value := range permissions[r] {
		if value == permission {
			return true
		}
	}

	return false
}
