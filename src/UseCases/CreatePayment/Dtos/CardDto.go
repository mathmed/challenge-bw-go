package Dtos

type CardDto struct {
	Cvv int `json:"cvv" binding:"required" validate:"required,numeric"`
	Number string `json:"number" binding:"required" validate:"required,numeric"`
	Expiration string `json:"expiration" binding:"required" validate:"required"`
	Flag string `json:"flag" binding:"required" validate:"required"`
	Owner string `json:"owner" binding:"required" validate:"required,alpha"`
	Parcels string `json:"parcels" binding:"required" validate:"required,numeric"`
	Amount float64 `json:"amount" binding:"required" validate:"required,numeric"`
}
