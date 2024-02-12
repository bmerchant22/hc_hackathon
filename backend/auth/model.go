package auth

import (
	"time"

	"github.com/bmerchant22/hc_hackathon/constants"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID       string         `gorm:"uniqueIndex" json:"user_id"` // roll or PF number
	Username     string         `json:"username"`
	Password     string         `json:"password"`
	RoleID       constants.Role `json:"role_id" gorm:"default:1"` // student role by default
	ExpiryDate   time.Time      `json:"expiry_date"`
}

type OTP struct {
	gorm.Model
	UserID  string `gorm:"column:user_id"`
	OTP     string `gorm:"column:otp"`
	Expires uint   `gorm:"column:expires"`
}