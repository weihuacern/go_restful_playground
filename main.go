package main

import (
	APIController "./controllers"
	"./db"
	"./middleware"
	"./utils"
	//"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	_ "strings"
)

func main() {
	log.Println("Starting server..")

	db.Init()

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		login := v1.Group("/login")
		{
			login.POST("/", APIController.LoginUser)
		}
		users := v1.Group("/users")
		{
			users.GET("/", APIController.GetUsers)
			users.POST("/", APIController.CreateUser)
			users.PUT("/:id", APIController.UpdateUser)
			users.DELETE("/:id", APIController.DeleteUser)
		}
		tasks := v1.Group("/tasks")
		{
			tasks.GET("/", APIController.GetTasks)
			tasks.POST("/", APIController.CreateTask)
			tasks.PUT("/:id", APIController.UpdateTask)
			tasks.DELETE("/:id", APIController.DeleteTask)
		}
	}

	//django api server from hua
	v0_django := r.Group("/api/v0")
	{
		v0_django.GET("/ds/servers", middleware.AuthMiddleWare(), func(c *gin.Context) {
			c.Redirect(http.StatusTemporaryRedirect, utils.GetEnv("PROXY_DATASOURCE_DJANGO_API", utils.PROXY_DATASOURCE_DJANGO_API)+c.Request.URL.String())
		})
		v0_django.GET("/ds/services", middleware.AuthMiddleWare(), func(c *gin.Context) {
			c.Redirect(http.StatusTemporaryRedirect, utils.GetEnv("PROXY_DATASOURCE_DJANGO_API", utils.PROXY_DATASOURCE_DJANGO_API)+c.Request.URL.String())
		})
	}
	//FIXME, /api/datashare/v1? need to redirect to somewhere
	v1_nodejs := r.Group("/api/v1")
	{
		v1_nodejs.GET("/contract", func(c *gin.Context) {
			c.Redirect(http.StatusTemporaryRedirect, utils.GetEnv("PROXY_DATASHARE_NODEJS_API", utils.PROXY_DATASHARE_NODEJS_API)+c.Request.URL.String())
		})
	}
	r.Run()
}
