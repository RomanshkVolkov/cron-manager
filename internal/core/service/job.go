package service

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/RomanshkVolkov/test-api/internal/adapters/repository"
	"github.com/RomanshkVolkov/test-api/internal/core/domain"
	"golang.org/x/net/html"
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

func SyncBoatyHealthChecker() domain.APIResponse[any] {
	client := &http.Client{}
	url := "https://boaty.com.mx"
	res, err := client.Get(url)
	if err != nil {
		return domain.APIResponse[any]{
			Success: false,
			Message: domain.Message{
				En: "Boaty health check failed",
				Es: "Comprobación de estado de Boaty fallida",
			},
		}
	}
	defer res.Body.Close()

	doc, err := html.Parse(res.Body)
	if err != nil {
		return domain.APIResponse[any]{
			Success: false,
			Message: domain.Message{
				En: "Boaty health check failed",
				Es: "Comprobación de estado de Boaty fallida",
			},
		}
	}

	isHealthy := false

	var crawler func(*html.Node)
	crawler = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "h1" {
			// print h1
			fmt.Println(n.LastChild.Data)
			content := n.LastChild.Data
			if strings.Contains(content, "YATE") {
				isHealthy = true
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			crawler(c)
		}
	}
	crawler(doc)

	if isHealthy {
		return domain.APIResponse[any]{
			Success: true,
			Message: domain.Message{
				En: "Boaty is healthy",
				Es: "Boaty está saludable",
			},
		}
	}
	BoatyHealingEvent()
	return domain.APIResponse[any]{
		Success: false,
		Message: domain.Message{
			En: "Boaty is unhealthy",
			Es: "Boaty está enfermo",
		},
	}
}

func BoatyHealingEvent() domain.APIResponse[any] {
	client := &http.Client{}
	url := "https://boaty.com.mx/clear-site"
	response, err := client.Get(url)

	faildResponse := domain.APIResponse[any]{
		Success: false,
		Message: domain.Message{
			En: "Boaty healing event failed",
			Es: "Evento de sanación de Boaty fallido",
		},
	}

	if err != nil {
		return faildResponse
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		repository.NewBoatyHealingRecord("New Healing Event")
		return domain.APIResponse[any]{
			Success: true,
			Message: domain.Message{
				En: "Boaty healing event succeeded",
				Es: "Evento de sanación de Boaty exitoso",
			},
		}
	}

	return faildResponse
}

// repository.NewBoatyHealingRecord("Healing Event")
