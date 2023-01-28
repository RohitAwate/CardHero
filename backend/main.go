package main

import "CardHero/api"

func main() {
	server := api.NewServer(8080)
	server.Start()
}
