package handler

import (
	"goPockemonParsingService/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *Handler) GetPockemons(c *gin.Context) {
	log.Info().Msg("from handler sent request to service: get all pockemons")
	pockemons, err := h.service.Product.GetPockemons()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dtos.GetPockemonsResponse{
		Data: pockemons,
	})
}


