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
	SlotTime  time.Time
	Status    string
}
