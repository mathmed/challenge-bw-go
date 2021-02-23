package repositories

import (
	"time"

	database "github.com/mathmed/challenge-bw-go/src/Infra/Database"
	models "github.com/mathmed/challenge-bw-go/src/Infra/Database/Models"
	dtos "github.com/mathmed/challenge-bw-go/src/UseCases/CreatePayment/Dtos"
)

// SaveDevice .
func SaveDevice(deviceData dtos.Device) uint {
	device := models.Device{
		IP: deviceData.IP,
		Device: deviceData.Device,
		Platform: deviceData.Platform,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	database.Instance.Create(&device)
	return device.ID
}