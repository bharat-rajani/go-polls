package api

import (
	"fmt"
	"github.com/bharat-rajani/go-polls/pkg/jcustom"
	"net/http"
	_ "net/http/pprof"
)

func RegisterRootRoutes(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("/", RootHandler)
	return mux
}

func RootHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {

	case "GET":
		switch request.URL.Path {
		case "/polls":
			GetPollsHandler(writer, request)
			return
		case "/hello":
			_, err := writer.Write([]byte(`HelloResponse`))
			if err != nil {
				return
			}
		default:
			writer.WriteHeader(http.StatusNotFound)
			fmt.Fprint(writer, "404 not found")
		}
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func GetPollsHandler(w http.ResponseWriter, r *http.Request) {
	b, err := jcustom.MarshalMap(map[string]interface{}{
		"1": map[string]interface{}{
			"name":   "Hello",
			"choice": []string{"hi", "bonjour"},
		},
	})
	if err != nil {
		w.Write([]byte("Internal Server Error"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(b)
	return
}
