package handlers

import (
	//"context"
	//"fmt"
	"fmt"
	"net/http"
	"strconv"

	//"time"

	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/LSD-Learn-Strive-Develop/logarithm/tree/main/rest-api/internal/models"
	//"github.com/9Neechan/Startup/internal/models"
	"github.com/gin-gonic/gin"
)

// POST task
func CreateTaskHandler(c *gin.Context) {
	var task_model models.Task

	if err := c.ShouldBindJSON(&task_model); err != nil {
		//c.String(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTask, err := task_model.CreateTask()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": createdTask})
}

// GET task by id
func GetTaskByIdHandler(c *gin.Context) {
	var task_model models.Task

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	foundTask, err := task_model.GetTaskById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, foundTask)
}

// DELETE task by id
func DeleteTaskByIdHandler(c *gin.Context) {
	var task_model models.Task

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = task_model.DeleteTaskById(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusOK)
}

func ChangeTaskStatusHandler(c *gin.Context) {
	var task_model models.Task

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = task_model.ChangeTaskStatus(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusOK)

}

// GET allTasks
func GetAllTasksHandler(c *gin.Context) {
	var task_model models.Task

	allTasks, err := task_model.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, allTasks)
}

func GetTasksByTgUserIdHandler(c *gin.Context) {
	var task_model models.Task

	tg_user_id, err := strconv.Atoi(c.Params.ByName("tg_user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userTasks, err := task_model.GetTasksByTgUserId(tg_user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userTasks)
}

func GetShiftTasksByTimeStampHandler(c *gin.Context) {
	fmt.Print("ellow")
	var task_model models.Task

	shift, err := strconv.Atoi(c.Params.ByName("shift"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	counter, err := strconv.Atoi(c.Params.ByName("counter"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shiftTasks, err := task_model.GetShiftTasksByTimeStamp(shift, counter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, shiftTasks)
}

func GetLastPlanByTgUserIdHandler(c *gin.Context) {
	var task_model models.Task

	tg_user_id, err := strconv.Atoi(c.Params.ByName("tg_user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	planTasks, err := task_model.GetLastPlanByTgUserId(tg_user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, planTasks)
}

/*
func (ts *taskServer) dueHandler(c *gin.Context) {
	badRequestError := func() {
		c.String(http.StatusBadRequest, "expect /due/<year>/<month>/<day>, got %v", c.FullPath())
	}

	year, err := strconv.Atoi(c.Params.ByName("year"))
	if err != nil {
		badRequestError()
		return
	}

	month, err := strconv.Atoi(c.Params.ByName("month"))
	if err != nil || month < int(time.January) || month > int(time.December) {
		badRequestError()
		return
	}

	day, err := strconv.Atoi(c.Params.ByName("day"))
	if err != nil {
		badRequestError()
		return
	}

	tasks := ts.store.GetTasksByDueDate(year, time.Month(month), day)
	c.JSON(http.StatusOK, tasks)
}
*/
