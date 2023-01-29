package models

import (
	"errors"
	"fmt"

	utils "github.com/Raajheer1/tamuhack-2023/api/m/v2/utils"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"size:255; not null;" json:"email"`
	Name     string `gorm:"size:255; not null;" json:"name"`
	Password string `gorm:"size:255; not null;" json:"password"`
}

func (u *User) CreateUser() (*User, error) {
	fmt.Println("model.CreateUser: Creating new user")
	fmt.Println(u)
	// Check that username does not already exist
	exists := DB.Where("name = ?", u.Name).First(&u)
	if !errors.Is(exists.Error, gorm.ErrRecordNotFound) {
		return &User{}, errors.New("user with that handle exists")
	}

	if err := DB.Create(&u).Error; err != nil {
		return &User{}, err
	}

	return u, nil
}

func Login(email string, password string) (*User, error) {
	u := User{}

	err := DB.Where("email = ?", email).Take(&u).Error
	if err != nil {
		return nil, err
	}
	if !utils.CheckPasswordHash(password, u.Password) {
		return nil, errors.New("incorrect password")
	}

	// userToken, err := utils.GenerateToken(u.ID)
	// if err != nil {
	// 	return nil, err
	// }

	return &u, nil
}

func GetUser(id uint) (*User, error) {
	fmt.Println("models.GetUser: getting user")
	var user User
	err := DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
