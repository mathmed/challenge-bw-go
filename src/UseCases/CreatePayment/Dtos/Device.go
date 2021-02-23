package dtos

// Device .
type Device struct {
	IP string `json:"ip" binding:"required" validate:"required,ip"`
	Device string `json:"device" binding:"required" validate:"required"`
	Platform string `json:"platform" binding:"required" validate:"required"`
}
