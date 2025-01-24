package service

import (
	"fmt"

	"github.com/RomanshkVolkov/test-api/internal/adapters/repository"
	"github.com/RomanshkVolkov/test-api/internal/core/domain"
)

func SyncElevaZapier() domain.APIResponse[any] {
	currentData := repository.GetContaboData()
	lenData := len(currentData)

	response := domain.APIResponse[any]{
		Success: false,
		Message: domain.Message{
			En: "Eleva Zapier sync failed",
			Es: "Sincronización de Eleva Zapier fallida",
		},
	}

	if lenData == 0 {
		response.Message.En = "No data to sync"
		response.Message.Es = "No hay datos para sincronizar"
		return response
	}

	rowAffects := repository.InsertProductionData(currentData)
	if rowAffects > 0 {
		countMessage := ": " + fmt.Sprint(rowAffects) + " - " + fmt.Sprint(lenData)
		repository.CleanContaboTable()

		response.Success = true
		response.Message.En = "Eleva Zapier sync done" + countMessage
		response.Message.Es = "Sincronización de Eleva Zapier completada" + countMessage

		return response
	}

	return response
}
