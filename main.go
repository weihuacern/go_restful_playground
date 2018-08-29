package main

import (
	APIController "./controllers"
	"./db"
	"github.com/gin-gonic/gin"
	"log"
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
	r.Run()
}
