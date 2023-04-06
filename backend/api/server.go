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
	router.Route("/api", func(router chi.Router) {
		router.Route("/{username:[a-z1-2_]+}", func(router chi.Router) {
			router.Route("/logs", func(router chi.Router) {
				router.Get("/", handlers.GetCards)
				router.Post("/", handlers.AddCard)
			})

			router.Route("/folders", func(router chi.Router) {
				router.Get("/*", handlers.GetFolders)
			})

			router.Route("/cards", func(router chi.Router) {
				router.Get("/*", handlers.GetCardsFromFolderPath)
			})

			router.Get("/search", handlers.Search)
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
