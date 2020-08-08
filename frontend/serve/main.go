package main

import (
	"net/http"

	"github.com/short-d/app/fw/env"
	"github.com/short-d/app/fw/envconfig"
	"github.com/short-d/app/fw/router"
	"github.com/short-d/app/fw/service"
	"github.com/short-d/short/frontend/serve/handle"
	"github.com/short-d/short/frontend/serve/shortapi"
	"github.com/short-d/short/frontend/serve/ssr"
)

func main() {
	goDotEnv := env.NewGoDotEnv()
	envConfig := envconfig.NewEnvConfig(goDotEnv)

	config := struct {
		ShortRPCAPIHost     string `env:"SHORT_RPC_API_HOST" default:"localhost"`
		ShortRPCAPIPort     int    `env:"SHORT_RPC_API_PORT" default:"8081"`
		HTTPPort    	    int    `env:"HTTP_PORT" default:"3000"`
	}{}

	err := envConfig.ParseConfigFromEnv(&config)
	if err != nil {
		panic(err)
	}

	gRPC, err := shortapi.NewGRPC(config.ShortRPCAPIHost, config.ShortRPCAPIPort)
	if err != nil {
		panic(err)
	}

	rootDir := "public"
	redirectPage := ssr.NewRedirectPage(rootDir)
	routes := []router.Route{
		{
			Method: http.MethodGet,
			Path:   "/r/:alias",
			Handle: handle.Redirect(redirectPage, gRPC),
		},
		{
			Method:      http.MethodGet,
			Path:        "/",
			MatchPrefix: true,
			Handle:      handle.File(rootDir),
		},
	}
	routingService := service.
		NewRoutingBuilder("Short").
		Routes(routes).
		Build()
	routingService.StartAndWait(config.HTTPPort)
}
