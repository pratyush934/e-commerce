package models

import (
	"github.com/pratyush934/e-commerce/user-service/dbUser"
	"github.com/rs/zerolog/log"
)

type Role struct {
	Id          int    `gorm:"primaryKey" json:"id"`
	RoleName    string `gorm:"not null; type:varchar(100)" json:"roleName"`
	Description string `gorm:"type:varchar(100)" json:"description"`
}

func CreateRoles(role *Role) error {
	err := dbUser.DB.Create(&role).Error

	if err != nil {
		log.Error().Msg("There is an issue while creating Roles in CreateRole/user-role.go")
		return err
	}

	return nil
}

func GetRoles() error {
	var roles []Role

	err := dbUser.DB.Find(&roles).Error

	if err != nil {
		log.Error().Msg("There is an issue while Getting Roles in GetRole/user-role.go")
		return err
	}
	return nil
}

func GetRoleById(roleId int) error {
	var role Role

	err := dbUser.DB.Where("id=?", roleId).First(&role).Error

	if err != nil {
		log.Error().Msg("There is an error in GetRoleById in GetRoleById/user-role.go")
		return err
	}
	return nil
}

func UpdateRole(role *Role) error {
	err := dbUser.DB.Updates(role).Error
	if err != nil {
		log.Error().Msg("There is na error in UpdateRole in UpdateRoleById/user-role.go")
		return err
	}
	return nil
}
