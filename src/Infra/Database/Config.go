package database

import (
	"os"

	models "github.com/mathmed/challenge-bw-go/src/Infra/Database/Models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Instance .
var Instance *gorm.DB

// Config .
func Config() {

	dsn := 
		os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + 
		"@tcp(" +os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT")+ ")/" +
		os.Getenv("DB_DATABASE")

	var err error
	Instance, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	Instance.AutoMigrate(
		models.PaymentStatus{},
		models.User{},
		models.Device{},
		models.Payment{},
	)

}