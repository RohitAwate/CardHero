package api

import (
	"CardHero/api/handlers"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type Server struct {
	Port   uint
	router http.Handler
}

func NewServer(port uint) Server {
	router := chi.NewRouter()
	router.Route("/api", func(api chi.Router) {
		api.Route("/{username:[a-z1-2_]+}", func(username chi.Router) {
			username.Get("/logs", handlers.GetUserLogs)
		})
	})

	return Server{
		Port:   port,
		router: router,
	}
}

func (s Server) Start() {
	addr := fmt.Sprintf(":%d", s.Port)
	err := http.ListenAndServe(addr, s.router)
	if err != nil {
		log.Fatalln(err)
	}
}
