package helpers

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/ritankarsaha/Golang-JWT-Auth/database"
	jwt "github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/joho/godotenv"
)

type SignedDetails struct {
	Email       string
	FirstName   string
	LastName    string
	Uid         string
	UserType    string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var secretKey []byte

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	secretKey = []byte(os.Getenv("SECRET_KEY"))
}

func GenerateAllTokens(email, firstname, lastname, uid, userType string) (signedToken, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Email:     email,
		FirstName: firstname,
		LastName:  lastname,
		Uid:       uid,
		UserType:  userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 24).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 168).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secretKey)
	if err != nil {
		log.Println("Error signing token:", err)
		return
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(secretKey)
	if err != nil {
		log.Println("Error signing refresh token:", err)
		return
	}

	return token, refreshToken, nil
}

func UpdateAllTokens(signedToken, signedRefreshToken, userId string) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	updateObj := bson.D{
		{Key: "token", Value: signedToken},
		{Key: "refresh_token", Value: signedRefreshToken},
		{Key: "updated_at", Value: time.Now().Format(time.RFC3339)},
	}

	filter := bson.M{"user_id": userId}
	upsert := true
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	result, err := userCollection.UpdateOne(
		ctx,
		filter,
		bson.D{{Key: "$set", Value: updateObj}},
		&opt,
	)

	if err != nil {
		log.Printf("Failed to update tokens for user %s: %v", userId, err)
		return
	}

	log.Printf("Updated tokens for user %s: %v", userId, result)
}