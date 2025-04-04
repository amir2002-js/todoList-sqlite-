package Handler

import "net/http"

type Todos struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Time  string `json:"time"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

	case http.MethodPost:

	case http.MethodPut:

	case http.MethodDelete:
	}
}
