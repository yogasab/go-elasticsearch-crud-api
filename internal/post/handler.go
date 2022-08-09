package post

import (
	"encoding/json"
	"log"
	"net/http"
)

type postHandler struct {
	postService PostService
}

func NewPostHandler(postService PostService) postHandler {
	return postHandler{postService: postService}
}

func (h postHandler) InsertDocument(w http.ResponseWriter, r *http.Request) {
	var req InsertDocumentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	result, err := h.postService.InsertDocument(r.Context(), req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	body, _ := json.Marshal(result)
	w.Write(body)
}
