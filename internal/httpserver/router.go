package httpserver

import (
	"time"

	"github.com/go-chi/chi"
	chimw "github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"github.com/payfazz/chrome-remote-debug/internal/httpserver/handlers/statement"
	"github.com/payfazz/chrome-remote-debug/internal/httpserver/handlers/status"
)

func (hs *HTTPServer) compileRouter() chi.Router {
	r := chi.NewRouter()

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Access-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	// A good base middleware stack
	r.Use(chimw.RequestID)
	r.Use(chimw.RealIP)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(chimw.Timeout(60 * time.Second))

	// create endpoint for server http request
	r.Method("GET", "/status", status.GetHandler())
	r.Method("POST", "/statement", statement.GetHandler())
	return r
}
