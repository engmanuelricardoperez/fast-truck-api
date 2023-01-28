package entrypoints

import "net/http"

type Handler interface {
	Handler(w http.ResponseWriter, r *http.Request) error
}
