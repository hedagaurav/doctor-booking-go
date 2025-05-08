package controllers

import (
	"doctor-booking-go/config"
	"doctor-booking-go/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// AddSlot handles POST /slots for doctors to create available slots
func AddSlot(c *gin.Context) {
	var input struct {
		Day       string `json:"day" binding:"required"`        // e.g. "Monday"
		StartTime string `json:"start_time" binding:"required"` // e.g. "10:00"
		EndTime   string `json:"end_time" binding:"required"`   // e.g. "13:00"
		Duration  int    `json:"duration" binding:"required"`   // in minutes
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the doctor ID from context (set by JWT middleware)
	doctorIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	doctorID := doctorIDVal.(uint)

	startTime, err1 := time.Parse("15:04", input.StartTime)
	endTime, err2 := time.Parse("15:04", input.EndTime)
	if err1 != nil || err2 != nil || endTime.Before(startTime) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start or end time"})
		return
	}

	slot := models.DoctorSlot{
		DoctorID:  doctorID,
		Day:       input.Day,
		StartTime: startTime,
		EndTime:   endTime,
		Duration:  input.Duration,
	}

	db := config.ConnectDB()
	if err := db.Create(&slot).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create slot"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Slot created", "slot": slot})
}

// GetDoctorSlots handles GET /slots for patients to view slots
func GetDoctorSlots(c *gin.Context) {
	doctorIDStr := c.Query("doctor_id")
	if doctorIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "doctor_id is required"})
		return
	}

	doctorID, err := strconv.ParseUint(doctorIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid doctor_id"})
		return
	}

	var slots []models.DoctorSlot
	db := config.ConnectDB()
	if err := db.Where("doctor_id = ?", doctorID).Find(&slots).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch slots"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"slots": slots})
}
