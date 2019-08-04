package modern

import (
	"context"
	"fmt"
	"net/http"
	"short/fw"
)

type HttpServer struct {
	mux    *http.ServeMux
	server *http.Server
	tracer fw.Tracer
	logger fw.Logger
}

func (s *HttpServer) ListenAndServe(port int) error {
	addr := fmt.Sprintf(":%d", port)

	s.server = &http.Server{Addr: addr, Handler: s.mux}
	err := s.server.ListenAndServe()

	if err == nil || err == http.ErrServerClosed {
		return nil
	}

	return err
}

func (s HttpServer) Shutdown() error {
	return s.server.Shutdown(context.Background())
}

func (s HttpServer) HandleFunc(pattern string, handler http.Handler) {
	s.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		w = setupPreFlight(w)
		if (*r).Method == "OPTIONS" {
			return
		}

		w = enableCors(w)
		r.Body = tap(r.Body, func(body string) {
			s.logger.Info(fmt.Sprintf("HTTP: url=%s host=%s method=%s body=%s", r.URL, r.Host, r.Method, body))
		})
		handler.ServeHTTP(w, r)
	})
}

func setupPreFlight(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	return w
}

func enableCors(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	return w
}

func NewHttpServer(logger fw.Logger, tracer fw.Tracer) HttpServer {
	mux := http.NewServeMux()

	return HttpServer{
		mux:    mux,
		tracer: tracer,
		logger: logger,
	}
}
