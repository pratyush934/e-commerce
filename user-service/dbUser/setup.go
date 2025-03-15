package dbUser

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() *gorm.DB {
	DB = connectDB()

	return DB
}

func connectDB() *gorm.DB {
	url := "root:Pratyush@123@/ecommerce-golang?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		log.Error().Msgf("There is an Error exist in dbUser/setup.go connectDB while parsing url %v", err.Error())
		return nil
	} else {
		log.Info().Msg("DB is successfully set")
	}
	return db
}
