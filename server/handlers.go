package server

import (
	"fmt"
	"github.com/codechaitu/circleci-heroku/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"os"
	"strconv"
)

func getDB(c *gin.Context) *gorm.DB {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err !=nil{
		fmt.Println("error in getDB: ", err)
	}
	return db
}

func hello(c *gin.Context) {
	fmt.Println("user taxi")
	c.String(200, "User, new change, good")
}

func getFromDB(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	db := getDB(c)
	var taxi models.Taxi
	db.Where("taxinum = ?", id).Find(&taxi)
	c.JSON(200, taxi)
}