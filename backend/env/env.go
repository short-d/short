package env

import "github.com/short-d/app/fw"

const (
	Production  fw.ServerEnv = "production"
	Staging     fw.ServerEnv = "staging"
	Development fw.ServerEnv = "development"
	Testing     fw.ServerEnv = "testing"
)
