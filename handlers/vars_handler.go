package handlers

import (
	//"context"
	//"fmt"
	"net/http"

	//"time"

	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/LSD-Learn-Strive-Develop/logarithm/tree/main/rest-api/internal/models"
	//"github.com/9Neechan/Startup/internal/models"
	"github.com/gin-gonic/gin"
)

func GetVarsHandler(c *gin.Context) {
	var vars_model models.Vars

	vars, err := vars_model.GetVars()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, vars)
}