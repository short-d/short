package risk

type Detector struct {
	blacklist BlackList
}

func (r Detector) IsURLMalicious(url string) bool {
	hasURL, err := r.blacklist.HasURL(url)
	if err != nil {
		return false
	}
	return hasURL
}

func NewDetector(blacklist BlackList) Detector {
	return Detector{blacklist: blacklist}
}
