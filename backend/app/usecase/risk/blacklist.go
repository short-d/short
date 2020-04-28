package risk

type URLBlackList interface {
	IsURLExist(url string) (bool, error)
}
