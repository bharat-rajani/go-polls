package internal

//
//import (
//	"github.com/rs/zerolog/log"
//	"io"
//	"net"
//	"net/http"
//	"time"
//)
//
//func S(){
//mux := http.NewServeMux()
//mux = RegisterRoutes(mux)
//	srv := http.Server{
//		Addr:              "localhost:8080",
//		Handler:           mux,
//		ReadTimeout:       5 * time.Second,
//		WriteTimeout:      5 * time.Second,
//		IdleTimeout:       30 * time.Minute,
//		ReadHeaderTimeout: 5 * time.Second,
//		ConnState: func(conn net.Conn, state http.ConnState) {
//			log.Info().Str("connection", conn.RemoteAddr().String()).Str("state", state.String()).Send()
//		},
//	}
//	err := srv.ListenAndServe()
//}
//func RegisterRoutes(mux *http.ServeMux) *http.ServeMux {
//	mux.HandleFunc("/hello", rootHandler)
//	return mux
//}
//
//func rootHandler(w http.ResponseWriter, r *http.Request) {
//	_, err := io.ReadAll(r.Body)
//	if err != nil {
//		return
//	}
//
//	mp := map[string]interface{}{
//		"test": map[string]string{
//			"nested": "hello",
//		},
//	}

//respBytes, err := jcustom.MarshalMap(mp)
//if err != nil {
//w.WriteHeader(http.StatusInternalServerError)
//return
//}
//_, err = w.Write(respBytes)
//if err != nil {
//return
//}
//}
