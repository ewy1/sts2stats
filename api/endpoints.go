//go:build api

package api

import (
	"encoding/json"
	"net/http"
)

func Init() {
	http.HandleFunc("/cards", ToJson(CardChoices))
	http.ListenAndServe(":6060", nil)
}

type HttpHandler = func() (any, error)

func ToJson(handler HttpHandler) http.HandlerFunc {
	f := func(w http.ResponseWriter, r *http.Request) {
		res, err := handler()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		data, err := json.Marshal(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		_, err = w.Write(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	return f
}
