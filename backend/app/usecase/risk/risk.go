package risk

// Detector determines whether the given items are malicious.
type Detector struct {
	blacklist BlackList
}

// IsShortLinkMalicious checks whether the given ShortLink is malicious.
func (r Detector) IsShortLinkMalicious(shortLink string) bool {
	hasShortLink, err := r.blacklist.HasShortLink(shortLink)
	if err != nil {
		return false
	}
	return hasShortLink
}

// NewDetector creates a new Detector
func NewDetector(blacklist BlackList) Detector {
	return Detector{blacklist: blacklist}
}
