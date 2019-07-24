package main

import (
	"tinyURL/dep"
	"tinyURL/modern"
)

func main() {
	service := dep.InitGraphQlService("TinyUrl GraphQL API", modern.GraphQlPath("/graphql"))
	service.Start(8080)

	service = dep.InitRoutingService("TinyUrl Routing API")
	service.StartAndWait(80)
}
