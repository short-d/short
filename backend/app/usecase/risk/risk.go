package risk

// Detector determines whether the given items are malicious.
type Detector struct {
	blacklist BlackList
}

// IsURLMalicious checks whether the given URL is malicious.
func (r Detector) IsURLMalicious(url string) bool {
	hasURL, err := r.blacklist.HasURL(url)
	if err != nil {
		return false
	}
	return hasURL
}

// NewDetector creates a new Detector
func NewDetector(blacklist BlackList) Detector {
	return Detector{blacklist: blacklist}
}
