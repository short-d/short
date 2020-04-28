package risk

type Detector struct {
	urlBlackList URLBlackList
}

func (r Detector) IsURLMalicious(url string) bool {
	isExist, err := r.urlBlackList.IsBlacklisted(url)
	if err != nil {
		return false
	}
	return isExist
}

func NewDetector(urlBlackList URLBlackList) Detector {
	return Detector{urlBlackList: urlBlackList}
}
