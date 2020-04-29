package permission

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
