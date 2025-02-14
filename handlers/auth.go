package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/vermakmanish001/otp-auth-app/config"
	"github.com/vermakmanish001/otp-auth-app/models"
	"github.com/vermakmanish001/otp-auth-app/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var usersCollection *mongo.Collection

func init() {
	config.ConnectDB()
	usersCollection = config.DB.Collection("users")
}

// user registeration
func Register(c *gin.Context) {
	var req struct {
		Phone    string `json:"phone"`
		DeviceID string `json:"device_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	//otp
	otp := utils.GenerateOTP()

	// saving otp in databse
	newUser := models.User{
		Phone:     req.Phone,
		OTP:       otp,
		ExpiresAt: time.Now().Add(5 * time.Minute), //expire in 5minutes
		DeviceID:  req.DeviceID,
	}

	_, err := usersCollection.InsertOne(context.TODO(), newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register"})
		return
	}

	fmt.Println("OTP for", req.Phone, "is", otp)

	c.JSON(http.StatusOK, gin.H{"message": "OTP sent"})
}

// login
func Login(c *gin.Context) {
	var req struct {
		Phone string `json:"phone"`
		OTP   string `json:"otp"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// user by otp and number
	var user models.User
	err := usersCollection.FindOne(context.TODO(), bson.M{"phone": req.Phone, "otp": req.OTP}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid OTP"})
		return
	}

	// checking if otp is expired or not
	if time.Now().After(user.ExpiresAt) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "OTP expired"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "phone": user.Phone})
}
