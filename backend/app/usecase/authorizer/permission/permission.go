package permission

// Permission is an approval of a mode of access to a resource
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
