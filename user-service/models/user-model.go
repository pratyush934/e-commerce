package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id             uuid.UUID   `gorm:"primaryKey; type:char(36)" json:"id"`
	Name           string      `gorm:"not null; type:varchar(100)" json:"name"`
	Email          string      `gorm:"unique;not null; type:varchar(100)" json:"email"`
	UserName       string      `gorm:"unique;not null; type:varchar(100)" json:"username"`
	PassWord       string      `gorm:"not null; type:varchar(100)" json:"passWord"`
	CreatedAt      time.Time   `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt      time.Time   `gorm:"autoUpdateTime" json:"updatedAt"`
	OrderId        []uuid.UUID `json:"orderId"`
	Address        []Address   `gorm:"foreignKey:UserId" json:"address"`
	PrimaryAddress uuid.UUID   `gorm:"type:char(36)" json:"primaryAddress"`
	RoleId         int         `gorm:"not null;index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"roleId"`
	Role           Role        `gorm:"foreignKey:RoleId;references:Id" json:"role"`
}
