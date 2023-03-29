package sorcery

import (
	"encoding/json"
	"net/http"
)

type Sorc func(*http.Request) (statusCode int, data interface{})

func (h Sorc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	statusCode, data := h(r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

var Http404handler http.Handler = Sorc(func(r *http.Request) (statusCode int, data interface{}) {
	statusCode = 404
	data = NewStatusMsg(404, "Not Found")
	return
})
