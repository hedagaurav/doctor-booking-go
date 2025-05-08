// File: models/appointment.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	DoctorID  uint
	PatientID uint
	// SlotDate is the date of the appointment
	// SlotTime is the time of the appointment
	SlotDate string
	SlotTime time.Time
	Status   string
}
