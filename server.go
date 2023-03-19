package main

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"smartway-test-task/internal/handler"
	"smartway-test-task/internal/repository"
	"smartway-test-task/internal/service"
	"time"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initialization config: %s", err.Error())
	}
	// TODO правильно я все сделал с паролем?
	pg, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
		Reload:   true,
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(pg)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	start(handlers.InitRoutes())
}

func start(r *mux.Router) {
	logrus.Printf("start application")
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initialization config: %s", err.Error())
	}

	serverAddress := fmt.Sprintf("%s:%s",
		viper.GetString("address"),
		viper.GetString("port"))

	srv := &http.Server{
		Handler:      r,
		Addr:         serverAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logrus.Printf("server is listening port %s", serverAddress)
	logrus.Fatal(srv.ListenAndServe())
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
