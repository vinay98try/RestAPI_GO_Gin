package controller

import (
	"PracticeGDB/config"
	"PracticeGDB/models"

	"github.com/gin-gonic/gin"
)

// function for getting all data/users details from postgreSQL
func GetUser(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}

// function for getting (GetOnlyUser) data/details of particular user from postgreSQL
func GetOnlyUser(c *gin.Context) {
	users := []models.User{}
	config.DB.Where("id = ?", c.Param("id")).Find(&users)
	c.JSON(200, gin.H{
		"status": "ok",
		"user":   &users,
	})
}

// function for creating new data/users details in postgreSQL
func CreateUser(c *gin.Context) {
	var user models.User
	c.ShouldBindJSON(&user)
	config.DB.Create(&user)
	c.JSON(200, &user)
}

// function for updating all data/details of particular user in postgreSQL
func UpdateUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).First(&user) //	db.First(&user)
	c.ShouldBindJSON(&user)                               // user.Name = "jinzhu 2", user.Age = 100
	config.DB.Save(&user)                                 // db.Save(&user)
	c.JSON(200, &user)
}

// function for deleting all data/details of particular user in postgreSQL
func DeleteUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(200, &user)
}
