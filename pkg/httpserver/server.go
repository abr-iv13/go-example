package httpserver

import (
	"time"

	"github.com/valyala/fasthttp"
)

const (
	_defaultReadTimeout        = 5 * time.Second
	_defaultWriteTimeout       = 5 * time.Second
	_defaultShutdownTimeout    = 3 * time.Second
	_defaultMaxRequestBodySize = 1024 * 1024 * 1024
)

type Server struct {
	server *fasthttp.Server
	notify chan error
	// shutdownTimeout time.Duration
}

func New(port string, handler fasthttp.RequestHandler, opts ...Option) *Server {
	httpServer := &fasthttp.Server{
		ReadTimeout:        _defaultReadTimeout,
		WriteTimeout:       _defaultWriteTimeout,
		MaxRequestBodySize: _defaultMaxRequestBodySize,
		Handler:            handler,
	}

	s := &Server{
		server: httpServer,
		notify: make(chan error, 1),
	}

	for _, opt := range opts {
		opt(s)
	}

	s.start(port)

	return s
}

func (s *Server) start(port string) {
	go func() {
		s.notify <- s.server.ListenAndServe(port)
		close(s.notify)
	}()
}

// Notify -.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *Server) Shutdown() error {
	return s.server.Shutdown()
}
