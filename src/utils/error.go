package utils

import (
	"log"
	"net/http"
)

func InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Internal Server Error: Method %v path: %v, error: %v", r.Method, r.URL.Path, err.Error())
	WriteJSONError(w, http.StatusInternalServerError, "Internal Server Error")
}

func BadRequest(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Bad Request Error: Method %v path: %v, error: %v", r.Method, r.URL.Path, err.Error())
	WriteJSONError(w, http.StatusBadRequest, "Bad Request Error")
}

func NotFound(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Not Found Error: Method %v path: %v, error: %v", r.Method, r.URL.Path, err.Error())
	WriteJSONError(w, http.StatusNotFound, "Not Found Error")
}
