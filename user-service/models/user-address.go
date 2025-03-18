package models

import (
	"github.com/google/uuid"
	"github.com/pratyush934/e-commerce/user-service/dbUser"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Address struct {
	Id       uuid.UUID `json:"address-id" gorm:"primaryKey; type:char(36)"`
	UserId   uuid.UUID `json:"userId" gorm:"type:char(36);index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Street   string    `json:"street"`
	Pin      int       `json:"pin"`
	City     string    `json:"city"`
	State    string    `json:"state"`
	LandMark string    `json:"landMark"`
}

func (address *Address) BeforeSave(db *gorm.DB) (err error) {
	address.Id = uuid.New()
	return
}

func (address *Address) SaveAddress() error {
	err := dbUser.DB.Create(address).Error
	if err != nil {
		log.Error().Msg("There is an issue in saving the address in SaveAddress/user-address.go")
		return err
	}
	return nil
}

func GetAddressByUserId(userId string) ([]Address, error) {
	var address []Address
	err := dbUser.DB.Where("userId=?", userId).Find(address).Error

	if err != nil {
		log.Error().Msg("There is an issue in GetAddressByUserId/user-address.go")
		return address, err
	}
	return address, nil
}

func GetAddressByPin(pin int) ([]Address, error) {
	var address []Address

	err := dbUser.DB.Where("pin=?", pin).Find(address).Error

	if err != nil {
		log.Error().Msg("There is an issue in GetAddressByPin/user-address.go")
		return address, err
	}
	return address, nil
}

func GetUserIdByPin(pin int) (string, error) {
	var userId string

	err := dbUser.DB.Where("pin=?", pin).Find(userId).Error

	if err != nil {
		log.Error().Msg("There is an issue in GetUsersByPin/user-address.go")
		return "", err
	}
	return userId, nil
}
