package server

import (
	"fmt"
	"github.com/codechaitu/circleci-heroku/models"
	"github.com/codechaitu/circleci-heroku/constants"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

func getDB(c *gin.Context) *gorm.DB {
	return c.MustGet(constants.ContextDB).(*gorm.DB)
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
	db.Where("taxi_no = ?", id).Find(&taxi)

	fmt.Println(taxi)
	c.JSON(200, taxi)
}