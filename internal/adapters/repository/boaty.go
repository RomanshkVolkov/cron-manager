package repository

import "github.com/RomanshkVolkov/test-api/internal/core/domain"

func NewBoatyHealingRecord(content string) {
	contaboDB := GetDBConnection("contabo")

	contaboDB.DB.AutoMigrate(&domain.BoatyHealingRecord{})

	contaboDB.DB.Create(&domain.BoatyHealingRecord{
		Content: content,
	})
}
