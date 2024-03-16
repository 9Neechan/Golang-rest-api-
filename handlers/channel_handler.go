package handlers

import (
	//"context"
	//"fmt"
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

func GetChannelsByTgUserIdHandler(c *gin.Context) {
	var ch_model models.Channel

	tg_user_id, err := strconv.Atoi(c.Params.ByName("tg_user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userTasks, err := ch_model.GetChannelsByTgUserId(tg_user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userTasks)
}