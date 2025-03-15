package models

type Role struct {
	Id          int    `gorm:"primaryKey" json:"id"`
	RoleName    string `gorm:"not null; type:varchar(100)" json:"roleName"`
	Description string `gorm:"type:varchar(100)" json:"description"`
}
