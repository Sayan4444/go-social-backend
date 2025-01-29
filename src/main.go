package main

import (
	"log"
	"net/http"

	"github.com/Sayan4444/go-social-backend/src/database"
	"github.com/Sayan4444/go-social-backend/src/models"
	"github.com/Sayan4444/go-social-backend/src/routes"
	"github.com/Sayan4444/go-social-backend/src/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	database.Connect()

	db := database.GetDB()
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Post{})

	r := mux.NewRouter().PathPrefix("/api").Subrouter()

	routes.RegisterPostRoutes(r)
	// http.Handle("/", r)
	log.Printf("Server started")
	port := utils.GetEnvAsString("PORT")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
