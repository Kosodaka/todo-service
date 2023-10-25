package app

import (
	"github.com/Kosodaka/todo-service/internal/repository"
	"github.com/Kosodaka/todo-service/internal/service"
	"github.com/Kosodaka/todo-service/internal/transport/rest"
	"github.com/Kosodaka/todo-service/internal/transport/rest/handler"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"github.com/uptrace/bun"
	"os"
)

var Db *bun.DB

func Run() {
	// Write logs in file
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	logger := zerolog.New(file).With().Timestamp().Logger()
	logger.Info().Msg("logging to a file")
	//  Initialize config
	if err = InitConfig(); err != nil {
		logger.Fatal().Msgf("error initializing configs: %s", err.Error())
	}
	// Load env variable
	if err = godotenv.Load(); err != nil {
		logger.Fatal().Msgf("error to load env variables: %s", err.Error())
	}
	// Initialize Postgres connection
	Db, err = repository.NewPostgresDb(repository.Config{
		Dsn: os.Getenv("DSN"),
	})
	if err != nil {
		logger.Fatal().Msgf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(Db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	// Start server
	srv := new(rest.Server)
	if err = srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logger.Fatal().Msg("error while running http server")
	}

}
func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
