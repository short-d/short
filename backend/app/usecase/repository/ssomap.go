package repository

// SSOMap accesses account mapping between SSOUser and internal User
// from storage media, such as database.
type SSOMap interface {
	IsSSOUserExist(ssoUserID string) (bool, error)
	GetShortUserID(ssoUserID string) (string, error)
	CreateMapping(sshUserID string, shortUserID string) error
}
