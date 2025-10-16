package grace

import (
	"context"
	"net/http"
	"os"
	"time"
)

const (
	k8sDefaultTerminationGracePeriod = 30 * time.Second
)

type (
	HTTPServer interface {
		ListenAndServe() error
		Shutdown(ctx context.Context) error
	}

	Server struct {
		shutdown <-chan os.Signal
		delegate HTTPServer
		timeout  time.Duration
	}

	ServerOption func(server *Server)
)

func WithShutdownSignal(ch <-chan os.Signal) ServerOption {
	return func(server *Server) {
		server.shutdown = ch
	}
}

func WithTimeout(timeout time.Duration) ServerOption {
	return func(server *Server) {
		server.timeout = timeout
	}
}

func NewServer(server HTTPServer, options ...ServerOption) *Server {
	srv := &Server{
		delegate: server,
		timeout:  k8sDefaultTerminationGracePeriod,
		shutdown: newInterruptSignalChannel(),
	}

	for _, option := range options {
		option(srv)
	}

	return srv
}

func (s *Server) ListenAndServe(ctx context.Context) error {
	select {
	case err := <-s.delegateListenAndServe():
		return err
	case <-ctx.Done():
		return s.shutdownDelegate(ctx)
	case <-s.shutdown:
		return s.shutdownDelegate(ctx)
	}
}

func (s *Server) delegateListenAndServe() chan error {
	listenErrCh := make(chan error)

	go func() {
		if err := s.delegate.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			listenErrCh <- err
		}
	}()

	return listenErrCh
}

func (s *Server) shutdownDelegate(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	if err := s.delegate.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		return err
	}
	return ctx.Err()
}
