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
		users := v1.Group("/login")
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
		nodes := v1.Group("/nodes")
		{
			appportals := nodes.Group("/appportals")
			{
				appportals.GET("/", APIController.GetAppPortals)
				appportals.POST("/", APIController.CreateAppPortal)
				appportals.PUT("/:id", APIController.UpdateAppPortal)
				appportals.DELETE("/:id", APIController.DeleteAppPortal)
			}
		}
	}

	r.Run()
}
