package controllers

import (
	 "context"
	 "github.com/ritankarsaha/Golang-JWT-Auth/database"
	 helper "github.com/ritankarsaha/Golang-JWT-Auth/helpers"
	 "github.com/ritankarsaha/Golang-JWT-Auth/models"
	 "log"
	 "net/http"
	 "time"
	 "github.com/gin-gonic/gin"
	 "github.com/go-playground/validator/v10"
	 "golang.org/x/crypto/bcrypt"
	 "go.mongodb.org/mongo-driver/bson"
	 "go.mongodb.org/mongo-driver/bson/primitive"
	 "go.mongodb.org/mongo-driver/mongo"

)

var userCollection = *&mongo.Collection = database.OpenCollection(database.Client,"user")
var validate = validator.New()

func HashPassowrd()