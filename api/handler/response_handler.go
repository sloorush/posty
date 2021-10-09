package handler

import (
	"encoding/json"
	"net/http"
	"posty/api/presenter"
)

func NewErrorResponse(status int, message string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	data := presenter.Error{Status: status, Message: message}
	json.NewEncoder(w).Encode(data)
}

func NewSuccessResponse(status int, message string, data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	resData := presenter.Success{Status: status, Message: message, Data: data}
	json.NewEncoder(w).Encode(resData)
}
