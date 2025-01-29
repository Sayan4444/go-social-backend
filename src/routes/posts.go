package routes

import (
	"github.com/Sayan4444/go-social-backend/src/controllers"
	"github.com/gorilla/mux"
)

func RegisterPostRoutes(r *mux.Router) {
	Posts := controllers.Posts{}
	postsRouter := r.PathPrefix("/posts").Subrouter()

	postsRouter.HandleFunc("", Posts.CreatePost).Methods("POST")
	postsRouter.HandleFunc("/{id}", Posts.GetPostByID).Methods("GET")
}
