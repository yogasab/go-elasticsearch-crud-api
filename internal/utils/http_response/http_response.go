package http_response

import (
	"encoding/json"
	"net/http"
	"yogasab/go-elasticsearch-crud-api/internal/utils/http_errors"
)

func NewJSONResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

func NewJSONResponseError(w http.ResponseWriter, err http_errors.RestErrors) {
	NewJSONResponse(w, err.Code(), err)
}
