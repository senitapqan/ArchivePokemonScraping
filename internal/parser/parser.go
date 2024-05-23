package parser

import "goPockemonParsingService/models"

type Parser interface {
	GetPageCount() (int, error)
	GetPockemonsOnPage(page int) ([]models.Pockemon, error)
}

type parser struct {}

func NewParser() Parser {
	return &parser{}
}