package repository

import (
	"github.com/RomanshkVolkov/test-api/internal/core/domain"
)

func GetContaboData() []domain.Zapier {
	contaboDB := GetDBConnection("contabo")

	zapierData := []domain.Zapier{}
	contaboDB.DB.Model(&domain.Zapier{}).Find(&zapierData)

	return zapierData
}

func CleanContaboTable() {
	contaboDB := GetDBConnection("contabo")
	contaboDB.DB.Exec("DELETE FROM zapier")
}

func InsertProductionData(data []domain.Zapier) int {
	productionDB := GetDBConnection("production")
	rowInserted := 0

	for _, zapier := range data {
		zapier.WhatsApp = SerializeAlphanumericString(zapier.WhatsApp)[:15]
		zapier.PhoneSocialMedia = SerializeAlphanumericString(zapier.PhoneSocialMedia)[:15]

		productionDB.DB.Create(&zapier)
		rowInserted++
	}

	return rowInserted
}
