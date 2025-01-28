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

	database.Connect()
	db := database.GetDB()
	db.AutoMigrate(&models.Post{})
	db.AutoMigrate(&models.User{})

	r := mux.NewRouter()

	routes.RegisterPostRoutes(r)
	http.Handle("/api", r)
	log.Printf("Server started")
	port := utils.GetEnvAsString("PORT")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
