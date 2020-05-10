package permission

// Permission represents access to a given resource.
type Permission int

const (
	CreateShortLink Permission = iota
	ViewShortLink
	EditShortLink
	DisableShortLink
	DeleteShortLink

	CreateChange
	ViewChange
	EditChange
	DeleteChange

	UpgradeUser
	DowngradeUser
	DisableUser
	DeleteUser

	ViewDashboards
)
