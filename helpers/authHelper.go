package helpers

import(
	"errors"
	"github.com/gin-gonic/gin"
)

//checking the user type before any fucntion is done primarily

func CheckUserType(c *gin.Context, roll string) (err error) {

	//checking the user type of the user over here.

	err = nil
	userType := c.GetString("user_type")
	if userType!=roll {
		err = errors.New("user type of the user cannot be matched by the databsase")
	}
	return err


}

func MathUserTypeToUid(c *gin.Context , userId string) (err error){
	userType := c.GetString("user_type")
	id := c.GetString("uid")
	err = nil
	if userType == "USER" && id!=userId{
		return errors.New("user id or uid cannot be properly matched for the user")
	}
	err = CheckUserType(c,userType)
	if err !=nil{
		return err
	}
}