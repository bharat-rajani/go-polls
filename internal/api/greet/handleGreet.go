package greet

import (
	"net/http"

	"github.com/bharat-rajani/go-polls/internal/service"
)

func HandleGreeter(s *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Log.Info().Msg("received request for greet ")
		_, err := w.Write([]byte(`HelloResponse`))
		if err != nil {
			return
		}
	}
}
