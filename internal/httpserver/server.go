package httpserver

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// HTTPServer represents the http server that handles the requests
type HTTPServer struct {
}

// Serve serves for http requests
func (hs *HTTPServer) Serve(port string) {
	r := hs.compileRouter()

	log.Printf("About to listen on %s. Go to http://127.0.0.1:%s", port, port)
	srv := http.Server{Addr: ":" + port, Handler: r}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalln("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

// NewServer ..
func NewServer() *HTTPServer {
	return &HTTPServer{}
}
