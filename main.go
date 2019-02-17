package main

import "circleci/server"

func main() {
	router := server.CreateRouter()
	server.StartServer(router)
}