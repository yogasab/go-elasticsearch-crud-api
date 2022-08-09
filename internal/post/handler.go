package post

import (
	"encoding/json"
	"net/http"
	"yogasab/go-elasticsearch-crud-api/internal/utils/http_response"
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
		data := map[string]interface{}{
			"status":  "failed",
			"code":    http.StatusBadRequest,
			"message": "failed to decode request body",
			"error":   err,
		}
		http_response.NewJSONResponse(w, http.StatusBadRequest, data)
		return
	}
	result, err := h.postService.InsertDocument(r.Context(), req)
	if err != nil {
		http_response.NewJSONResponseError(w, err)
		return
	}
	http_response.NewJSONResponse(w, http.StatusCreated, result)
}
