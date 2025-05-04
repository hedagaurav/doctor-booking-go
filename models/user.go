// File: models/user.go
package models

import (
	"gorm.io/gorm"
)

// User represents a user in the system
// It can be a doctor or a patient
// The Role field indicates the type of user
// "doctor" for doctors and "patient" for patients
type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Role     string // "doctor", "patient"
}
