package risk

// Detector uses a blacklist to determine if a URL
// is considered safe and allowed.
type Detector struct {
	blacklist BlackList
}

// IsURLMalicious verifies if the blacklist contains
// the given url string
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
