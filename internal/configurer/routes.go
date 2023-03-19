package configurer

import (
	"github.com/bharat-rajani/go-polls/internal/api"
	"github.com/bharat-rajani/go-polls/internal/api/polls"
	"github.com/bharat-rajani/go-polls/internal/api/votes"
	"github.com/bharat-rajani/go-polls/internal/service"
	"net/http"
	"net/http/pprof"
)

func RegisterRoutes(svc *service.Service) {
	RegisterVoteRoutes(svc)
	RegisterPollsRoutes(svc)
	svc.RegisterRoute("/service/hello", api.HandleGreeter)
	svc.RegisterRoute("/debug/pprof/", func(service *service.Service) http.HandlerFunc {
		return pprof.Index
	})
	svc.RegisterRoute("/debug/pprof/cmdline", func(service *service.Service) http.HandlerFunc {
		return pprof.Cmdline
	})
	svc.RegisterRoute("/debug/pprof/profile", func(service *service.Service) http.HandlerFunc {
		return pprof.Profile
	})
	svc.RegisterRoute("/debug/pprof/symbol", func(service *service.Service) http.HandlerFunc {
		return pprof.Symbol
	})
	svc.RegisterRoute("/debug/pprof/trace", func(service *service.Service) http.HandlerFunc {
		return pprof.Trace
	})
}

func RegisterPollsRoutes(s *service.Service) {
	s.RegisterRoute("/api/vote", votes.HandleListVotes)
}
func RegisterVoteRoutes(s *service.Service) {
	s.RegisterRoute("/api/polls", polls.HandleListPolls)
}
