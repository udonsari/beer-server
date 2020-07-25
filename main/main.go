package main

import "github.com/UdonSari/beer-server/main/server"

func main() {
	server := server.New()
	server.Init()
	server.Start()
}
