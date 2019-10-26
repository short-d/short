package provider

import "short/app/adapter/kgs"

// KgsRPCConfig includes hostname and port for key generation service API
type KgsRPCConfig struct {
	Hostname string
	Port     int
}

// NewKgsRPC creates RPC
func NewKgsRPC(config KgsRPCConfig) (kgs.RPC, error) {
	return kgs.NewRPC(config.Hostname, config.Port)
}
