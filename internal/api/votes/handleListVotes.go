package votes

import (
	"github.com/bharat-rajani/go-polls/internal/service"
	"github.com/bharat-rajani/go-polls/pkg/jcustom"
	"net/http"
)

func HandleListVotes(s *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Log.Info().Msg("Inside vote")
		b, err := jcustom.MarshalMap(map[string]interface{}{
			"1": map[string]interface{}{
				"name":   "Hello",
				"choice": []string{"hi", "bonjour"},
			},
		})
		if err != nil {
			_, _ = w.Write([]byte("Internal server Error"))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, _ = w.Write(b)
	}
}
