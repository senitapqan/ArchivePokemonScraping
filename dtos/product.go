package dtos

import "goPockemonParsingService/models"

type GetPockemonsResponse struct {
	Data []models.Pockemon
}