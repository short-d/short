package fw

import (
	"context"
	"fmt"
	"net/http"
)

type Server interface {
	ListenAndServe(port int) error
	Shutdown() error
}

type HttpServer struct {
	handler http.Handler
	server  *http.Server
}

func (s *HttpServer) ListenAndServe(port int) error {
	addr := fmt.Sprintf(":%d", port)

	s.server = &http.Server{Addr: addr, Handler: s.handler}
	err := s.server.ListenAndServe()

	if err == nil || err == http.ErrServerClosed {
		return nil
	}

	return err
}

func (s HttpServer) Shutdown() error {
	return s.server.Shutdown(context.Background())
}

func NewHttpServer(handler http.Handler, logger Logger) HttpServer {
	mux := http.NewServeMux()

	mux.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		logger.Info(fmt.Sprintf("Request %v", r))
		handler.ServeHTTP(w, r)
	})

	return HttpServer{
		handler: mux,
	}
}
