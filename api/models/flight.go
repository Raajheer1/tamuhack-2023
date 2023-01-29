package models

import (
	"fmt"
	"github.com/Raajheer1/tamuhack-2023/api/m/v2/duffel"
	"github.com/jinzhu/gorm"
)

type Flight struct {
	gorm.Model
	User    User   `json:"user"`
	UserID  uint   `json:"user_id"`
	OfferID string `json:"offer_id"`
}

func CreateFlight(user *User, offerID string) error {
	fmt.Println("model.CreateFlight() Creating new flight")
	var flight Flight
	flight.User = *user
	flight.UserID = user.ID
	flight.OfferID = offerID

	if err := DB.Create(&flight).Error; err != nil {
		return err
	}

	return nil
}

func GetUserBookings(user *User) ([]duffel.Offer, error) {
	fmt.Println("models.GetUserBookings: getting user bookings")
	var flights []Flight
	err := DB.Where("user_id = ?", user.ID).Preload("User").Find(&flights).Error
	if err != nil {
		return nil, err
	}

	var bookings []duffel.Offer
	for _, flight := range flights {
		booking, err := duffel.GetSingleOfferById(flight.OfferID)
		if err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}

	return bookings, nil
}
