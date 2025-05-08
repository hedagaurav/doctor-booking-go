package cmd

import (
	"doctor-booking-go/config"
	"doctor-booking-go/models"
	"doctor-booking-go/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		panic("Error loading .env file")
	}
	db := config.ConnectDB()
	db.AutoMigrate(&models.User{}, &models.DoctorSlot{}, &models.Appointment{})

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
