package main

import (
	APIController "./src/controllers"
	"./src/db"
	"./src/middleware"
	"./src/utils"
	//"fmt"
	"github.com/gin-gonic/gin"
	"log"
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
		v0_django.GET("/ds/servers", middleware.AuthMiddleWare(utils.ROLEACC_DATASOURCE_DJANGO), APIController.ProxyRedirect("PROXY_DATASOURCE_DJANGO_API"))
		v0_django.GET("/ds/services", middleware.AuthMiddleWare(utils.ROLEACC_DATASOURCE_DJANGO), APIController.ProxyRedirect("PROXY_DATASOURCE_DJANGO_API"))
	}
	//FIXME, /api/datashare/v1? need to redirect to somewhere
	v1_nodejs := r.Group("/api/v1")
	{
		v1_nodejs.GET("/contract", middleware.AuthMiddleWare(utils.ROLEACC_DATASHARE_NODEJS), APIController.ProxyRedirect("PROXY_DATASHARE_NODEJS_API"))
	}
	r.Run()
}
