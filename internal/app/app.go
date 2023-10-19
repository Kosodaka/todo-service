package app

import (
	"github.com/Kosodaka/todo-service/internal/handler"
	"github.com/Kosodaka/todo-service/internal/repository"
	"github.com/Kosodaka/todo-service/internal/server"
	"github.com/Kosodaka/todo-service/internal/service"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"os"
)

func Run() {
	// Write logs in file
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	logger := zerolog.New(file).With().Timestamp().Logger()
	logger.Info().Msg("logging to a file")
	//  Initialize config
	if err = initConfig(); err != nil {
		logger.Fatal().Msgf("error initializing configs: %s", err.Error())
	}
	// Load env variable
	if err = godotenv.Load(); err != nil {
		logger.Fatal().Msgf("error to load env variables: %s", err.Error())
	}
	// Initialize Postgres connection
	db, err := repository.NewPostgresDb(repository.Config{
		Dsn: os.Getenv("DSN"),
	})
	if err != nil {
		logger.Fatal().Msgf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	// Start server
	srv := new(server.Server)
	if err = srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logger.Fatal().Msg("error while running http server")
	}
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
