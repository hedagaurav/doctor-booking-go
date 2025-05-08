package routes

import (
	"doctor-booking-go/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		doctor := api.Group("/doctor").Use(AuthMiddleware("doctor"))
		doctor.POST("/slots", controllers.AddSlot)

		patient := api.Group("/patient").Use(AuthMiddleware("patient"))
		patient.GET("/slots", GetAvailableSlots)
		patient.POST("/appointments", BookAppointment)

		api.GET("/appointments", AuthMiddleware("any"), GetAppointments)
	}
}
