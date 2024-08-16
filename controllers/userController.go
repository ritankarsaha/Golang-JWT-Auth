package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ritankarsaha/Golang-JWT-Auth/database"
	helper "github.com/ritankarsaha/Golang-JWT-Auth/helpers"
	"github.com/ritankarsaha/Golang-JWT-Auth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(bytes)
}

func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("user_id")

		if err := helper.MathUserTypeToUid(ctx, userId); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		var user models.User
		err := userCollection.FindOne(c, bson.M{"user_id": userId}).Decode(&user)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				ctx.JSON(http.StatusNotFound, gin.H{
					"error": "User not found",
				})
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
			}
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	}
}