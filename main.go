package main

import (
	"tinyURL/app"
	"tinyURL/dep"
	"tinyURL/modern"
)

func main() {
	service := dep.InitGraphQlService("TinyUrl GraphQL API", modern.GraphQlPath("/graphql"))
	service.Start(8080)

	service = dep.InitRoutingService("TinyUrl Routing API", app.WwwRoot("app/web/build"))
	service.StartAndWait(80)
}
