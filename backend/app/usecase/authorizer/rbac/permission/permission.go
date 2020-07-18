package permission

// Permission represents access to a given resource.
type Permission int

const (
	ViewAdminPanel Permission = iota
	CreateShortLink
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

	CreateAPIKey
)
