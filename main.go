package main

import "tinyURL/dep"

func main() {
	service := dep.InitGraphQlService("TinyUrl API")
	service.StartAndWait(8080)
}
