package models

import (
	"fmt"

	duffel "github.com/Raajheer1/tamuhack-2023/api/m/v2/duffel"
	"github.com/jinzhu/gorm"
)

type Flight struct {
	gorm.Model
	User   User         `json:"user"`
	UserID uint         `json:"user_id"`
	Offer  duffel.Offer `json:"offer"`
}

func CreateFlight(user User, offer duffel.Offer) error {
	fmt.Println("model.CreateFlight() Creating new flight")
	var flight Flight
	flight.User = user
	flight.UserID = user.ID
	flight.Offer = offer
	if err := DB.Create(&flight).Error; err != nil {
		return err
	}

	return nil
}
