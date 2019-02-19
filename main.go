package main

import (
	"github.com/jinzhu/gorm"
	"log"
	"os"

	"github.com/codechaitu/circleci-heroku/db"
	"github.com/codechaitu/circleci-heroku/server"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	database, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
		panic("failed to establish database connection")
	}
	defer database.Close()
	db.Init(database)

	router := server.CreateRouter(database)
	server.StartServer(router)
}