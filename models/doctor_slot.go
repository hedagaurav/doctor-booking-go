// File: models/doctor_slot.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type DoctorSlot struct {
	gorm.Model
	DoctorID  uint
	Day       string
	StartTime time.Time
	EndTime   time.Time
	// gap in between 2 sessions.
	// 0 means no gap
	// 5 means 5 minute gap

	GapBetweenSessions int
	Duration           int
}
