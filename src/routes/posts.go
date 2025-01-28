package routes

import (
	"github.com/Sayan4444/go-social-backend/src/controllers"
	"github.com/gorilla/mux"
)

func RegisterPostRoutes(router *mux.Router) {
	Posts := controllers.Posts{}
	router.HandleFunc("/posts", Posts.CreatePost).Methods("POST")
	router.HandleFunc("/posts/{id}", Posts.GetPostByID).Methods("GET")
}
