package api

import (
	"fmt"
	"log"
	"os"
	"pustaka-api/api/controllers"

	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func init() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Gagal membuka file konfigurasi")
	}
}

func Run() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error getting ENV")
	} else {
		fmt.Println("Get Values ENV")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	apiPort := fmt.Sprintf(":%s", os.Getenv("API_PORT"))
	fmt.Printf("Listening to port %s", apiPort)

	server.Run(apiPort)
}
