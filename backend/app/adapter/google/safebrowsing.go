package google

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/usecase/risk"
)

var _ risk.URLBlackList = (*SafeBrowsing)(nil)

const safeBrowsingLookupAPI = "https://safebrowsing.googleapis.com/v4/threatMatches:find"

type SafeBrowsing struct {
	apiKey      string
	httpRequest fw.HTTPRequest
}

func (s SafeBrowsing) IsBlacklisted(url string) (bool, error) {
	api := s.auth(safeBrowsingLookupAPI)
	body := lookupAPIRequest{
		ThreatInfo: threatInfo{
			ThreatTypes: []threatType{
				malWare,
				socialEngineering,
				potentiallyHarmfulApp,
				unwantedSoftware,
			},
			PlatformTypes: []platformType{
				allPlatforms,
			},
			ThreatEntries: []threat{
				{URL: url},
			},
		},
	}

	buf, err := json.Marshal(body)
	if err != nil {
		return false, err
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	res := lookupAPIResponse{}
	err = s.httpRequest.JSON(http.MethodPost, api, headers, string(buf), &res)
	if err != nil {
		return false, err
	}
	return len(res.Matches) > 0, nil
}

func (s SafeBrowsing) auth(baseURL string) string {
	return fmt.Sprintf("%s/?key=%s", baseURL, s.apiKey)
}

func NewSafeBrowsing(apiKey string, req fw.HTTPRequest) SafeBrowsing {
	return SafeBrowsing{
		apiKey:      apiKey,
		httpRequest: req,
	}
}

type lookupAPIRequest struct {
	ThreatInfo threatInfo `json:"threatInfo"`
}

type threatInfo struct {
	ThreatTypes   []threatType   `json:"threatTypes"`
	PlatformTypes []platformType `json:"platformTypes"`
	ThreatEntries []threat       `json:"threatEntries"`
}

type lookupAPIResponse struct {
	Matches []match `json:"matches"`
}

type match struct {
	ThreatType      threatType   `json:"threatType"`
	PlatformType    platformType `json:"platformType"`
	Threat          threat       `json:"threat"`
	CacheDuration   string       `json:"cacheDuration"`
	ThreatEntryType string       `json:"threatEntryType"`
}

type threatType string

const (
	malWare               threatType = "MALWARE"
	socialEngineering     threatType = "SOCIAL_ENGINEERING"
	potentiallyHarmfulApp threatType = "POTENTIALLY_HARMFUL_APPLICATION"
	unwantedSoftware      threatType = "UNWANTED_SOFTWARE"
)

type platformType string

const (
	allPlatforms platformType = "ALL_PLATFORMS"
	anyPlatform  platformType = "ANY_PLATFORM"
	window       platformType = "WINDOWS"
	linux        platformType = "LINUX"
	osx          platformType = "OSX"
	chrome       platformType = "CHROME"
	ios          platformType = "IOS"
	android      platformType = "Android"
)

type threat struct {
	URL string `json:"url"`
}
