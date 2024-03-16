package main

import (
	"fmt"
	"os"

	"github.com/LSD-Learn-Strive-Develop/logarithm/tree/main/rest-api/configs"
	"github.com/LSD-Learn-Strive-Develop/logarithm/tree/main/rest-api/handlers"

	//"github.com/9Neechan/Startup/configs"
	//"github.com/9Neechan/Startup/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.ConnectDB()

	router := gin.Default()

	taskRoutes := router.Group("/task")
	taskRoutes.POST("/", handlers.CreateTaskHandler)
	taskRoutes.GET("/:id", handlers.GetTaskByIdHandler)
	taskRoutes.GET("/tg_user_id/:tg_user_id", handlers.GetTasksByTgUserIdHandler)
	taskRoutes.DELETE("/:id", handlers.DeleteTaskByIdHandler)
	taskRoutes.GET("/all/", handlers.GetAllTasksHandler)
	taskRoutes.POST("/status/:id", handlers.ChangeTaskStatusHandler)
	taskRoutes.GET("/sorted/:shift/:counter", handlers.GetShiftTasksByTimeStampHandler)

	channelRoutes := router.Group("/channel")
	channelRoutes.GET("/tg_user_id/:tg_user_id", handlers.GetChannelsByTgUserIdHandler)

	varsRoutes := router.Group("/vars")
	varsRoutes.GET("/", handlers.GetVarsHandler)

	planRoutes := router.Group("/plan")
	planRoutes.GET("/last/:tg_user_id", handlers.GetLastPlanByTgUserIdHandler)

	router.Run("0.0.0.0:" + os.Getenv("SERVERPORT")) //
	fmt.Println("Server running on port ", os.Getenv("SERVERPORT"))

	//router.Run("pmpu.site/log_go:5007")
	//router.Run("pmpu.site/neechan_test:5002")
}

// https://pmpu.site/neechan_test/task
// {"task_id":50 , "title":"ahahah",  "tg_user_id":4646, "status": false, "creaste_time": 0, "change_status_time": 0, "channel_id": "neechan", "channel_username": "neech", "plan_id": 12}
