package models

import (
	"fmt"
	duffel "github.com/Raajheer1/tamuhack-2023/api/m/v2/duffel"
	"github.com/jinzhu/gorm"
	"hash/fnv"
	"reflect"
)

type Flight struct {
	gorm.Model
	User    User         `json:"user"`
	UserID  uint         `json:"user_id"`
	OfferID uint         `json:"offer_id"`
	Offer   duffel.Offer `json:"offer"`
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func CreateFlight(user *User, offer duffel.Offer) error {
	fmt.Println("model.CreateFlight() Creating new flight")
	var flight Flight
	flight.User = *user
	flight.UserID = user.ID
	//fmt.Println("Offer passed to model: ", offer)
	if err := DB.Create(&offer).Error; err != nil {
		fmt.Println("Error creating offer record: ", err)
		return err
	}

	var offer_record duffel.Offer
	if err := DB.Where("id = ?", offer.ID).First(&offer_record).Error; err != nil {
		fmt.Println("Error selecting offer record: ", err)
		return err
	}

	flight.Offer = offer_record
	hashedID := uint(hash(offer_record.ID))
	fmt.Println("Hashed ID: ", hashedID)
	fmt.Println(reflect.TypeOf(hashedID))
	flight.OfferID = hashedID

	if err := DB.Create(&flight).Error; err != nil {
		return err
	}

	return nil
}

func GetUserBookings(user *User) ([]Flight, error) {
	fmt.Println("models.GetUserBookings: getting user bookings")
	var flights []Flight
	err := DB.Where("user_id = ?", user.ID).Preload("User").Preload("Offer").Find(&flights).Error
	if err != nil {
		return nil, err
	}
	return flights, nil
}
