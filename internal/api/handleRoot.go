package api

//func (s *service.Service) handleRoot() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		s.Log.Info().Msg("request received")
//		switch r.Method {
//
//		case "GET":
//			switch r.URL.Path {
//			case "/":
//				w.Write([]byte(`Welcome to polls service server`))
//				return
//			default:
//				w.WriteHeader(http.StatusNotFound)
//				fmt.Fprint(w, "404 not found")
//			}
//		default:
//			w.WriteHeader(http.StatusMethodNotAllowed)
//		}
//	}
//}
