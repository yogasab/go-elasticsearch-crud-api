package post

import (
	"encoding/json"
	"net/http"
	"yogasab/go-elasticsearch-crud-api/internal/utils/http_response"

	"github.com/julienschmidt/httprouter"
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

func (h postHandler) FindDocumentByID(w http.ResponseWriter, r *http.Request) {
	id := httprouter.ParamsFromContext(r.Context()).ByName("id")

	post, err := h.postService.FindDocumentByID(r.Context(), id)
	if err != nil {
		data := map[string]interface{}{
			"status":  err.Status(),
			"code":    err.Code(),
			"message": err.Message(),
			"error":   err.Data(),
		}
		http_response.NewJSONResponse(w, http.StatusBadRequest, data)
		return
	}
	http_response.NewJSONResponse(w, http.StatusCreated, post)
}
