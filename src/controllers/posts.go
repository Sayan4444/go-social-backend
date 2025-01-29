package controllers

import (
	"net/http"

	"github.com/Sayan4444/go-social-backend/src/database"
	"github.com/Sayan4444/go-social-backend/src/models"
	"github.com/Sayan4444/go-social-backend/src/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type createPostPayload struct {
	Content string   `validate:"required,max=100"`
	Title   string   `validate:"required,max=100"`
	Tags    []string `json:"tags"`
}

type Posts struct{}

var validate *validator.Validate

func (p *Posts) CreatePost(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	var payload createPostPayload
	err := utils.ReadJSON(r, &payload)
	if err != nil {
		utils.BadRequest(w, r, err)
		return
	}

	validate = validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(payload)
	if err != nil {
		utils.BadRequest(w, r, err)
		return
	}

	post := models.Post{
		Content: payload.Content,
		Title:   payload.Title,
		Tags:    payload.Tags,
	}

	result := db.Create(&post)

	if result.Error != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, "Failed to create book")
		return
	}

	utils.WriteJSON(w, http.StatusCreated, post)
}

func (p *Posts) GetPostByID(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	vars := mux.Vars(r)
	id := vars["id"]

	var post models.Post
	result := db.First(&post, id)

	if result.Error != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, "Failed to retrieve books")
		return
	}
	utils.WriteJSON(w, http.StatusOK, post)
}
