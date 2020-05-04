package role

import "github.com/short-d/short/app/usecase/authorizer/permission"

// Role represents a groupings of permissions
type Role int

const (
	Basic Role = iota

	SecuritySpecialist

	ShortLinkViewer
	ShortLinkEditor

	ChangeLogViewer
	ChangeLogEditor

	Admin
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

// IsAllowed tells if the given role grants access to a permission
func (r Role) IsAllowed(permission permission.Permission) bool {
	for _, value := range permissions[r] {
		if value == permission {
			return true
		}
	}

	return false
}
