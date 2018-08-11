package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"github.com/gin-gonic/gin/binding"
	"../db"
	"../models"
)

func GetAppPortals(c *gin.Context) {

	var appportals []models.AppPortal
	db := db.GetDB()
	db.Find(&appportals)
	c.JSON(200, appportals)
}

func CreateAppPortal(c *gin.Context) {
	var appportal models.AppPortal
	var db = db.GetDB()

	if err := c.BindJSON(&appportal); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	db.Create(&appportal)
	c.JSON(http.StatusOK, &appportal)
}

func UpdateAppPortal(c *gin.Context) {
	id := c.Param("id")
	var appportal models.AppPortal

	db := db.GetDB()
	if err := db.Where("id = ?", id).First(&appportal).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.BindJSON(&appportal)
	db.Save(&appportal)
	c.JSON(http.StatusOK, &appportal)
}

func DeleteAppPortal(c *gin.Context) {
	id := c.Param("id")
	var appportal models.AppPortal
	db := db.GetDB()

	if err := db.Where("id = ?", id).First(&appportal).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	db.Delete(&appportal)
}
