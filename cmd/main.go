package main

import (
    "github.com/ibrahaknazarov/credit_risk/internal/configs"
    "github.com/ibrahaknazarov/credit_risk/internal/controller"
    "github.com/ibrahaknazarov/credit_risk/internal/db"
    "github.com/ibrahaknazarov/credit_risk/internal/repository"
    "github.com/ibrahaknazarov/credit_risk/internal/service"
)
// @title credit_risk system
// @contact.name credit_risk API Support
// @contact.url https://test.com/
// @contact.email iamhaknazarov@gmail.com
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	logger.Info().Msg("Starting up - Start")
	if err := configs.ReadSettings(); err != nil {
		logger.Error().Err(err).Msg("Error reading settings" + err.Error())
		return
	}

	dbConn, err := db.InitConnection()
	if err != nil {
		logger.Error().Err(err).Msg("Error during database connection initialization: " + err.Error())
		return
	}

	repo := repository.NewRepository(dbConn)
	svc := service.NewService(repo)
	ctrl := controller.NewController(svc)

	if err = ctrl.InitRoutes(); err != nil {
		logger.Error().Err(err).Msg("Error during http-service initialization: " + err.Error())
		return
	}

	if err = db.CloseConnection(dbConn); err != nil {
		logger.Error().Err(err).Msg("Error during database connection close: " + err.Error())
		return
	}

	logger.Info().Msg("Starting up - End")
}
