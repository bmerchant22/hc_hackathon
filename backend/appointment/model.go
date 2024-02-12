package appointment

import (
	"github.com/bmerchant22/hc_hackathon/auth"
	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	UserID string `json:"user_id"` // Foreign key referencing User.UserID
	User   auth.User   `gorm:"foreignKey:UserID;references:UserID"`
	Note string `json:"note"`
	DocId string `json:"doc_id"` // user id of the doctor
	StatusCode int `json:"status_code"`
	Prescription string `json:"prescription"` // url of the image of the prescription
}
