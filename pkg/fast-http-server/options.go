package fasthttpserver

import (
	"time"
)

// Option -.
type Option func(*Server)

// Port -.
// func Port(port string) Option {
// 	return func(s *Server) {
// 		s.server.Addr = net.JoinHostPort("", port)
// 	}
// }

// ReadTimeout -.
func ReadTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.ReadTimeout = timeout
	}
}

// WriteTimeout -.
func WriteTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.WriteTimeout = timeout
	}
}

// MaxRequestBodySize -.
func MaxRequestBodySize(maxRequestBodySize int) Option {
	return func(s *Server) {
		s.server.MaxRequestBodySize = maxRequestBodySize
	}
}
