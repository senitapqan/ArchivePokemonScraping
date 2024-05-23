package main

import (
	"goPockemonParsingService/internal/handler"
	"goPockemonParsingService/internal/parser"
	"goPockemonParsingService/internal/service"
	"goPockemonParsingService/server"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if err := initConfig(); err != nil {
		log.Fatal().Err(err).Msg("some error with initializiing")
	}

	parser := parser.NewParser()
	service := service.NewService(parser)
	handler := handler.NewHandler(service)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatal().Err(err).Msg("error with run server")
	}

}

func initConfig() error {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}