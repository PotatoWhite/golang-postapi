package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"postapi/app/models"
)

func parse(r *http.Request, dest interface{}) error {
	return json.NewDecoder(r.Body).Decode(dest)
}

func sendResponse(w http.ResponseWriter, r *http.Request, status int, body interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return
		}
		fmt.Fprint(w, &err, body)
	}
}

func mapPostToJSON(p *models.Post) models.JsonPost {
	return models.JsonPost{
		Id: p.Id,
		Title: p.Title,
		Content:  p.Content,
		Author: p.Author,
	}
}
