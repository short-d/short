package provider

import "short/app/adapter/kgs"

type KgsRpcConfig struct {
	Hostname string
	Port     int
}

func NewKgsRpc(config KgsRpcConfig) (kgs.Rpc, error) {
	return kgs.NewRpc(config.Hostname, config.Port)
}
