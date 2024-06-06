package server

import (
	"backend/internal/server/api"
	"backend/internal/server/db"
	utils "backend/utils/util"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Server struct {
	Address   string
	ApiServer *api.ApiServer
	DB        *db.ApiDatabase
}

func NewServer(addr string) error {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error al cargar el archivo .env")
	}
	apiserver := api.NewApiServer(addr)
		connectionString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"))
	db := db.NewApiDatabase(connectionString)
	storage, _ := db.Run()
	err := apiserver.Run(storage)
	utils.CheckError(err)
	return nil
}
