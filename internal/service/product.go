package service

import (
	"fmt"
	"goPockemonParsingService/internal/parser"
	"goPockemonParsingService/models"

	"github.com/rs/zerolog/log"
)

type ProductService struct {
	parser parser.Parser
}

func NewProductService(parser parser.Parser) *ProductService {
	return &ProductService{
		parser: parser,
	}
}

func (s *ProductService) GetPockemons() ([]models.Pockemon, error) {
	log.Info().Msg("from service sent request to parser: get page count")
	pageCount, err := s.parser.GetPageCount()
	if err != nil {
		return nil, fmt.Errorf("error on parsing page count with message: %s", err.Error())
	}

	log.Info().Msgf("whole count of pages: %d", pageCount)
	var pockemons []models.Pockemon

	for page := 1; page <= pageCount; page++ {
		log.Info().Msgf("from service sent request to parser: get pockemons from page #%d", page)
		pockemonsOnPage, err := s.parser.GetPockemonsOnPage(page)
		if err != nil {
			return nil, fmt.Errorf("error on pagsring page %d with message: %s", page, err.Error())
		}
		log.Info().Msgf("scrapped all pockemons from page #%d", page)
		pockemons = append(pockemons, pockemonsOnPage...)
	}

	return pockemons, nil
}