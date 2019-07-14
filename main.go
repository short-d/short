package main

import (
	"tinyURL/dep"
	"tinyURL/modern"
)

func main() {
	service := dep.InitGraphQlService("TinyUrl API", modern.GraphQlPath("/graphql"))
	service.StartAndWait(8080)
}
