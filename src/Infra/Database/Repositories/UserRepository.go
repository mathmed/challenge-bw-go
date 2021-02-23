package repositories

import (
	"strings"
	"time"

	database "github.com/mathmed/challenge-bw-go/src/Infra/Database"
	models "github.com/mathmed/challenge-bw-go/src/Infra/Database/Models"
	dtos "github.com/mathmed/challenge-bw-go/src/UseCases/CreatePayment/Dtos"
)

// SaveUser .
func SaveUser(userData dtos.User) uint {
	userBirthDate := strings.Split(userData.BirthDate, "/")
	user := models.User{
		Name: userData.Name,
		Email: userData.Email,
		Document: userData.Document,
		BirthDate: userBirthDate[2] + "/" + userBirthDate[1] + "/" + userBirthDate[0],
		Phone: userData.Phone,
		Street: userData.Address.Street,
		Number: userData.Address.Number,
		Neighborhood: userData.Address.Neighborhood,
		City: userData.Address.City,
		Area: userData.Address.AdministrativeArea,
		PostalCode: userData.Address.PostalCode,
		Complement: userData.Address.Complement,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	database.Instance.FirstOrCreate(&user)
	return user.ID
}