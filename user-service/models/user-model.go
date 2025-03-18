package models

import (
	"github.com/google/uuid"
	"github.com/pratyush934/e-commerce/user-service/dbUser"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html"
	"strings"
	"time"
)

type User struct {
	Id             uuid.UUID   `gorm:"primaryKey; type:char(36)" json:"id"`
	Name           string      `gorm:"not null; type:varchar(100)" json:"name"`
	Email          string      `gorm:"unique;not null; type:varchar(100)" json:"email"`
	UserName       string      `gorm:"unique;not null; type:varchar(100)" json:"username"`
	PassWord       string      `gorm:"not null; type:varchar(100)" json:"password"`
	CreatedAt      time.Time   `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt      time.Time   `gorm:"autoUpdateTime" json:"updatedAt"`
	OrderId        []uuid.UUID `json:"orderId"`
	Address        []Address   `gorm:"foreignKey:UserId" json:"address"`
	PrimaryAddress uuid.UUID   `gorm:"type:char(36)" json:"primaryAddress"`
	RoleId         int         `gorm:"not null;index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"roleId"`
	Role           Role        `gorm:"foreignKey:RoleId;references:Id" json:"role"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.Id = uuid.New()
	return
}

func (user *User) BeforeSave(db *gorm.DB) (err error) {
	password, err := bcrypt.GenerateFromPassword([]byte(user.PassWord), bcrypt.DefaultCost)

	if err != nil {
		log.Warn().Msg("There is an error in the method user-model.go/BeforeSave Part1")
		return
	}

	user.PassWord = string(password)
	user.UserName = html.EscapeString(strings.TrimSpace(user.UserName))

	return err
}

func (user *User) Save() (*User, error) {
	err := dbUser.DB.Create(user).Error

	if err != nil {
		log.Error().Msg("There is an error in user-model.go/Save")
		return nil, err
	}

	return user, nil
}

func (user *User) ValidatePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(password))

	if err != nil {
		log.Error().Msg("There is an Error in user-model.go/ValidatePassword")
		return err
	}

	return err
}

func GetUserById(userId string) (*User, error) {

	var user User

	err := dbUser.DB.Where("id=?", userId).First(&user).Error

	if err != nil {
		log.Error().Msg("There is an error in user-model.go/GetUserById")
		return &User{}, err
	}
	return &user, err
}

func GetUserByUserName(userName string) (*User, error) {
	var user User

	err := dbUser.DB.Where("username=?", userName).First(&user).Error

	if err != nil {
		log.Error().Msg("There is an error in user-model.go/GetUserByUserName")
		return &User{}, err
	}

	return &user, err
}

func GetUserByEmail(userEmail string) (*User, error) {
	var user User

	err := dbUser.DB.Where("email=?", userEmail).First(&user).Error
	if err != nil {
		log.Error().Msg("There is an error in user-model.go/GetUserByEmail")
		return &User{}, err
	}
	return &user, err
}

func GetUsers(user *[]User) error {

	err := dbUser.DB.Find(&user).Error

	if err != nil {
		log.Error().Msg("There is an error in user-model.go/GetUsers")
		return err
	}
	return nil
}

func UpdateUser(user *User) error {
	err := dbUser.DB.Omit("password").Updates(user).Error

	if err != nil {
		log.Error().Msg("There is an error in user-model.go/UpdateUser")
		return err
	}
	return nil
}
