package app

import (
	"fmt"
	"log"
	"net/http"
	"postapi/app/models"
)

func (a *App) IndexHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World")
	}
}

func (a *App) CreatePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.PostRequest{}
		err := parse(r, &req)
		if err != nil {
			log.Println("Cannot parse post body, err=", err.Error())
			sendResponse(w, r, http.StatusBadRequest, nil)
			return
		}

		p := &models.Post{
			Id:      0,
			Title:   req.Title,
			Content: req.Content,
			Author:  req.Author,
		}

		id, err := a.DB.CreatePost(p)
		if err != nil {
			log.Print("Could not save post in DB", err)
			sendResponse(w, r, http.StatusInternalServerError, r.Body)
			return
		}
		p.Id = id
		log.Print("Create post with id=", id)

		resp := mapPostToJSON(p)
		sendResponse(w, r, http.StatusCreated, resp)
	}
}
