package risk

type Detector struct {
	urlBlackList BlackList
}

func (r Detector) IsURLMalicious(url string) bool {
	isExist, err := r.urlBlackList.HasURL(url)
	if err != nil {
		return false
	}
	return isExist
}

func NewDetector(urlBlackList BlackList) Detector {
	return Detector{urlBlackList: urlBlackList}
}
