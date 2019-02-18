package main

import "circleci-heroku/server"

func main() {
	router := server.CreateRouter()
	server.StartServer(router)
}