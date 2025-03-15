package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	log.Info().Msg("I am Pratyush and I am building this project")

	r := gin.Default()

	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusAccepted, gin.H{
			"message": "Hello World",
		})
	})

	err := r.Run()

	if err != nil {
		log.Warn().Msg("Error in starting the server")
	} else {
		log.Info().Msg("Server is running successfully")
	}
}
