package helpers

import (
	"context"
	"log"
	"time"
	"github.com/ritankarsaha/Golang-JWT-Auth/database"
	jwt "github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct {
	Email       string
	First_Name   string
	Last_Name    string
	Uid         string
	UserType    string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client,"user")
var secretKey = []byte("secretkey")

func GenerateAllTokens(email string, firstname string, lastname string, uid string, userType string ) (signedToken string, signedRefreshToken string, err error){

	claims := &SignedDetails{
		Email : email,
		First_Name: firstname,
		Last_Name: lastname,
		Uid: uid,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},

	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err!=nil{
		log.Panic(err)
		return 
	}

	return token, refreshToken, err
	
}

