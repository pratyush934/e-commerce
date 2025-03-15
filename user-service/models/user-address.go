package models

import "github.com/google/uuid"

type Address struct {
	Id       uuid.UUID `json:"address-id" gorm:"primaryKey; type:char(36)"`
	UserId   uuid.UUID `json:"userId" gorm:"type:char(36);index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Street   string    `json:"street"`
	Pin      int       `json:"pin"`
	City     string    `json:"city"`
	State    string    `json:"state"`
	LandMark string    `json:"landMark"`
}
