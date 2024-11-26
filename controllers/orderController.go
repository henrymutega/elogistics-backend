package controllers

import (
	"context"
	"elogistics-backend/database"
	"elogistics-backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order.ID = primitive.NewObjectID()
	order.Status = "Pending"
	order.CreatedAt = primitive.NewDateTimeFromTime(time.Now())

	collection := database.Client.Database("elogistics").Collection("orders")
	_, err := collection.InsertOne(context.Background(), order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully"})
}
