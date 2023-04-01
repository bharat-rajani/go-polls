package greet

import (
	"github.com/bharat-rajani/go-polls/internal/service"
	"net/http"
)

func HandleGreeter(s *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`HelloResponse`))
		if err != nil {
			return
		}
	}
}
