package helper

import (
	"errors"

	"github.com/diegobermudez03/golang-jwt-auth/models"
)

func MatchUserTypeToUID(user models.User, userId string) error {
	if user.UserType == "USER" && user.ID != userId{
		err := errors.New("Unathorized to access this resource")
		return err 
	}
	return nil
}