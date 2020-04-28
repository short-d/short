package env

import "github.com/short-d/app/fw"

const (
	// Production implies the server is running under production environment.
	Production  fw.ServerEnv = "production"
	// Staging implies the server is running under staging environment.
	Staging     fw.ServerEnv = "staging"
	// Development implies the server is running on developers local machine.
	Development fw.ServerEnv = "development"
	// Testing implies the server is running under testing environment.
	Testing     fw.ServerEnv = "testing"
)
