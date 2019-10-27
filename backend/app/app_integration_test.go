// +build integration

package app

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/byliuyang/app/modern/mdservice"

	"short/dep"
	"short/dep/provider"
)

const (
	TestServerPort = "8081"
	BaseURI        = "http://127.0.0.1"
)

// TestIntegration_HealthURL ensures that the health endpoint exists and is
// returning an expected result.
func TestIntegration_HealthURL(t *testing.T) {
	_, client, closer := apiServer()
	defer closer()

	url := fmt.Sprintf("%s:%s/health", BaseURI, TestServerPort)
	res, err := client.Get(url)
	assertError(t, err, nil)
	assertStatus(t, res.StatusCode, http.StatusOK)
}

// assertStatus ensures that the provided http status matches expected.
func assertStatus(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("HTTP Status is not equal!\ngot:  %d\nwant: %d", got, want)
	}
}

// assertError ensures that the provided error matches expected.
func assertError(t *testing.T, got, want error) {
	if got != want {
		t.Errorf("Errors are not equal!\ngot:  %s\nwant: %s", got, want)
	}
}

// apiServer sets up and returns a test API server along with a http.Client for
// communicating to the test server.
func apiServer() (service mdservice.Service, client *http.Client, teardown func()) {
	// Bring up the short service.
	httpAPI := dep.InjectRoutingService(
		"Routing API",
		nil,
		provider.GithubClientID(""),
		provider.GithubClientSecret(""),
		provider.JwtSecret(""),
		provider.WebFrontendURL(""),
	)

	port, _ := strconv.Atoi(TestServerPort)
	httpAPI.Start(port)

	// Generate a HTTP Client for performing requests.
	client = &http.Client{
		Timeout: time.Second * 5,
	}

	return httpAPI, client, httpAPI.Stop
}
