package main

import "detroit-testing/app/infra"

func main() {
	server := infra.NewInfra(":9999")
	server.RunGin()
}
