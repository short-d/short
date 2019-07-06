package main

import "tinyURL/dep"

func main() {
	service := dep.InitGraphQlService("TinyUrl API")
	service.Start(8080)

	select {}
}
