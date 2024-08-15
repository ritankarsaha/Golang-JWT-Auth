package helpers

import (
	"errors"
	"github.com/gin-gonic/gin"
)

// CheckUserType checks the user type before any function is executed.
func CheckUserType(c *gin.Context, roll string) error {
	userType := c.GetString("user_type")
	if userType != roll {
		return errors.New("user type of the user cannot be matched by the database")
	}
	return nil
}

// MathUserTypeToUid matches user type to UID and checks for consistency.
func MathUserTypeToUid(c *gin.Context, userId string) error {
	userType := c.GetString("user_type")
	id := c.GetString("uid")

	if userType == "USER" && id != userId {
		return errors.New("user id or uid cannot be properly matched for the user")
	}

	err := CheckUserType(c, userType)
	if err != nil {
		return err
	}
	return nil
}