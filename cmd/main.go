package main

import (
	"log"
	"os"

	"github.com/Lunovoy/simple-api/internal/handler"
	"github.com/Lunovoy/simple-api/internal/repository"
	"github.com/joho/godotenv"
)

func main() {
	if err := initApp(); err != nil {
		log.Printf("error while initiating app: %s", err.Error())
		panic(err)
	}
	defer repository.CloseMongoDB()

	port := os.Getenv("PORT")

	app := handler.InitRoutes()

	app.Listen(":" + port)
}

func initApp() error {
	if err := loadENV(); err != nil {
		return err
	}

	if err := repository.StartMongoDB(); err != nil {
		return err
	}
	return nil
}

func loadENV() error {

	goENV := os.Getenv("GO_ENV")

	if goENV == "" {

		if err := godotenv.Load(); err != nil {
			return err
		}
	}
	return nil
}
