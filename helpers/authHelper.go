package helper

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) error {
	userType := c.GetString("user_type")
	var err error = nil

	if userType != role {
		msg := fmt.Sprintf("Unauthorized to access this resource: userType (%s) does not match the role (%s)", userType, role)
		err = errors.New(msg)
		return err
	}
	return err
}

func MatchUserTypeToUid(c *gin.Context, userId string) error {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	var err error = nil

	// this guarantees a user can only access their own data and no other user's data
	if userType == "USER" && uid != userId {
		err = errors.New("unauthorized to access this resource (uid != userId)")
		return err
	}

	err = CheckUserType(c, userType)
	if err != nil {
		return err
	}

	fmt.Println("Match userType to uid passed with no errors")
	return err

}
