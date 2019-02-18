package main

import(
	"github.com/codechaitu/circleci-heroku/server"
	"github.com/codechaitu/circleci-heroku/db"
	"github.com/codechaitu/circleci-heroku/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {

	router := server.CreateRouter()
	server.StartServer(router)
}