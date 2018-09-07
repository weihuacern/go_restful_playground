package api

import (
	"../db"
	"../models"
	"../utils"
	//"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {

	var users []models.User
	db := db.GetDB()
	db.Find(&users)
	c.JSON(200, users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	var db = db.GetDB()

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	db.Create(&user)
	c.JSON(http.StatusOK, &user)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	db := db.GetDB()
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.BindJSON(&user)
	db.Save(&user)
	c.JSON(http.StatusOK, &user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	db := db.GetDB()

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	db.Delete(&user)
}

func LoginUser(c *gin.Context) {
	var login models.User
	var db = db.GetDB()

	if err := c.BindJSON(&login); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var username string
	var password string
	var role string
	rows, err := db.Raw("select user_name, password, role from users where user_name = ? and password = ? and status = 'active'", login.UserName, login.Password).Rows()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()
	for rows.Next() {
		rows.Scan(&username, &password, &role)
	}
	if username == "" || password == "" || role == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not a valid user"})
		return
	}

	seed := utils.GetEnv("GLOBAL_SEED", utils.GLOBAL_SEED)
	key := []byte(seed)
	jwt_token, err := utils.GenJWTString(key, username, role)
	c.Header("Xauth", jwt_token)
	c.JSON(http.StatusOK, gin.H{"success": "Login success, Xauth is issued"})
}
