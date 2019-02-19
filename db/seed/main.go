package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/codechaitu/circleci-heroku/db"
	"github.com/codechaitu/circleci-heroku/models"

	"github.com/jinzhu/gorm"
 _ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	var taxis []models.Taxi

	getData("data/taxis.json", &taxis)
	database, err := gorm.Open("postgres", "postgres://czzunqplhnfsbr:5e7afbe690b47e5e10138fb64d68630e5e8f3dd77c672a2cfd78dd7ff0bdc521@ec2-107-20-167-11.compute-1.amazonaws.com:5432/d35alq6mfcq8or")
	if err != nil {
		panic("failed to establish database connection")
	}
	defer database.Close()
	db.Init(database)

	for _, taxi := range taxis {
		db.DB.Save(&taxi)
	}

}

func getData(fileName string, v interface{}) {
	file, _ := os.Open(fileName)
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, v)
}
