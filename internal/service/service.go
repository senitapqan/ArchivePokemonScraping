package service

import (
	"goPockemonParsingService/models"
	"goPockemonParsingService/internal/parser"
)

type Product interface {
	GetPockemons() ([]models.Pockemon, error)
}

type Service struct {
	Product
}

func NewService(parser parser.Parser) *Service {
	return &Service{
		Product: NewProductService(parser),
	}
}