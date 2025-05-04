package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "root:yourpassword@tcp(localhost:3306)/doctor_booking?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// handle error
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	// err = db.AutoMigrate(&models.User{}, &models.Appointment{}, &models.Doctor{})
	// if err != nil {
	// 	panic("failed to migrate database")
	// }
	// db.AutoMigrate(&models.User{}, &models.Appointment{}, &models.Doctor{})
	// if err != nil {
	// 	panic("failed to migrate database")
	// }
	return db
}
