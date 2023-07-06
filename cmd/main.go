package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sharifsharifzoda/project-management-system"
	"github.com/sharifsharifzoda/project-management-system/configs"
	"github.com/sharifsharifzoda/project-management-system/db"
	"github.com/sharifsharifzoda/project-management-system/logging"
	"github.com/sharifsharifzoda/project-management-system/pkg/handler"
	"github.com/sharifsharifzoda/project-management-system/pkg/repository"
	"github.com/sharifsharifzoda/project-management-system/pkg/service"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logging.InitLog()
	logger := logging.GetLogger()

	if err := godotenv.Load(); err != nil {
		logger.Fatal(err)
	}
	//reading from yaml
	if err := InitConfigs(); err != nil {
		logger.Fatalf("error while reading config file. error is %v", err.Error())
	}

	var cfg configs.DatabaseConnConfig

	if err := viper.Unmarshal(&cfg); err != nil {
		logger.Fatalf("Couldn't unmarshal the config into struct. error is %v", err.Error())
	}
	cfg.Password = os.Getenv("DB_PASSWORD")

	conn := db.GetDBConnection(cfg)

	db.Init(conn)

	//---------- Dependency injection-----------
	newRepository := repository.NewRepository(conn)
	newService := service.NewService(newRepository, logger)
	newHandler := handler.NewHandler(newService.Auth, newService.User, newService.Project, newService.Task)
	//--------------------------------------------

	server := new(project_management_system.Server)
	go func() {
		if err := server.Run(os.Getenv("PORT"), newHandler.InitRoutes()); err != nil {
			logger.Fatalf("error while running http.server. Error is %s", err.Error())
		}
	}()
	fmt.Printf("Server is listening to port: %s\n", os.Getenv("PORT"))

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	db.Close(conn)

	fmt.Println("server is shutting down")
	if err := server.Shutdown(context.Background()); err != nil {
		logger.Fatalf("error while shutting server down. Error: %s", err.Error())
	}
}

func InitConfigs() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	return viper.ReadInConfig()
}
