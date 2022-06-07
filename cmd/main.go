package main

import (
	"fmt"
	"os"

	"github.com/Muhammadjon226/api_gateway/api"
	"github.com/Muhammadjon226/api_gateway/config"
	"github.com/Muhammadjon226/api_gateway/pkg/event"
	"github.com/Muhammadjon226/api_gateway/pkg/logger"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/joho/godotenv"
)

var (
	log logger.Logger
	cfg config.Config
	err error
)

func initDeps() {
	cfg = config.Load()
	log = logger.New(cfg.LogLevel, "api_gateway")

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	_, err := gormadapter.NewAdapter("postgres", psqlString, true)
	if err != nil {
		log.Error("new adapter error", logger.Error(err))
		return
	}
}

func main() {
	fmt.Println("hello")

	if info, err := os.Stat(".env"); !os.IsNotExist(err) {
		if !info.IsDir() {
			godotenv.Load(".env")
		}
	}
	initDeps()

	var kafka *event.Kafka
	if cfg.Environment == "develop" {
		kafka, err = event.NewKafka(cfg, log)
		if err != nil {
			panic(err)
		}
	}

	server := api.New(api.Config{
		Logger: log,
		Config: cfg,
		Kafka:  kafka,
	})

	fmt.Println("port: ", cfg.HTTPPort)
	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("error while running gin server", logger.Error(err))
	}
}
