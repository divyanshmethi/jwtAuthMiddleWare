package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) error {
	userType := c.GetString("user_type")
	var err error
	if userType != role {
		err = errors.New("Unauthorized access to this resource")
	}
	return err
}

func MatchUserTypeToUid(c *gin.Context, userID string) error {
	loggedInUserType := c.GetString("user_type")
	logginInUserID := c.GetString("uid")
	if loggedInUserType == "ADMIN" || (loggedInUserType == "USER" && logginInUserID == userID) {
		return nil
	}
	return errors.New("Unauthorized to access this resource")
}
