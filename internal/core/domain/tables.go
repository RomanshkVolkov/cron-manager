package domain

import (
	"time"

	"gorm.io/gorm"
)

// table users - usuarios
type User struct {
	gorm.Model
	UserData
	ProfileID         uint      `gorm:"not null" json:"-"`
	OTP               string    `gorm:"type:nvarchar(6)" json:"otp"` // One Time Password
	OTPExpirationDate time.Time `gorm:"column:otp_expiration_date" json:"otpExpirationDate"`
	Password          string    `gorm:"type:nvarchar(200);not null" json:"-" validate:"required,min=6,max=200"`
}

type Dev struct {
	gorm.Model
	Tag string `gorm:"type:nvarchar(200);not null" json:"tag" validate:"required,min=3,max=200"`
	IP  string `gorm:"type:nvarchar(200);not null" json:"ip" validate:"required,min=3,max=200"`
}

type Zapier struct {
	Fecha            time.Time `gorm:"type:DATE;column:fecha" json:"fecha"`
	Nombre           string    `gorm:"type:VARCHAR(200);column:nombre" json:"nombre"`
	Apellido         string    `gorm:"type:VARCHAR(200);column:apellido" json:"apellido"`
	Email            string    `gorm:"type:VARCHAR(200);column:email" json:"email"`
	WhatsApp         string    `gorm:"type:VARCHAR(200);column:whatsapp" json:"whatsapp"`
	PhoneSocialMedia string    `gorm:"type:VARCHAR(200);column:phone_social_media" json:"phone_social_media"`
	Ciudad           string    `gorm:"type:VARCHAR(200);column:ciudad" json:"ciudad"`
	Development      int       `gorm:"type:INT;column:development" json:"development"`
	Comments         string    `gorm:"type:LONGTEXT;column:comments" json:"comments"`
	Presupuesto      string    `gorm:"type:VARCHAR(200);column:presupuesto" json:"presupuesto"`
	Interes          string    `gorm:"type:VARCHAR(200);column:interes" json:"interes"`
	Origen           string    `gorm:"type:VARCHAR(200);column:origen" json:"origen"`
	Platform         string    `gorm:"type:VARCHAR(200);column:platform" json:"platform"`
	Campanya         string    `gorm:"type:VARCHAR(200);column:campanya" json:"campanya"`
	AdSet            string    `gorm:"type:VARCHAR(200);column:adset" json:"adSet"`
	AdName           string    `gorm:"type:VARCHAR(200);column:adname" json:"adName"`
	Broker           string    `gorm:"type:VARCHAR(10);column:broker" json:"broker"`
	FechaCRM         time.Time `gorm:"type:DATE;column:fecha_crm" json:"fechaCRM"`
	InsertionDate    time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP;column:insertion_date" json:"insertionDate"`
}

type BoatyHealingRecord struct {
	gorm.Model
	Content string `gorm:"type:MEDIUMTEXT;" json:"content"`
}

type Zapier_bkp struct {
	Zapier
}
