package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/OlegBabakov/pow-server/config"
	"github.com/OlegBabakov/pow-server/internal/repository"
	"github.com/OlegBabakov/pow-server/pkg/logger"
	"github.com/OlegBabakov/pow-server/pkg/pow"
	"github.com/OlegBabakov/pow-server/pkg/pow/hashcash"
)

// Server structure.
type Server struct {
	cfg      *config.ServerConfig
	logger   logger.Logger
	verifier pow.Verifier
	repo     repository.Repositories
	listener net.Listener

	wg          sync.WaitGroup
	connections chan net.Conn
}

// InitWithConfig init server instance with config.
func InitWithConfig(cfg *config.ServerConfig, log logger.Logger) *Server {
	powProvider, err := hashcash.NewPOW(cfg.Pow.Complexity)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepositories()

	return newServer(cfg, log, powProvider, repo)
}

func newServer(
	cfg *config.ServerConfig,
	logger logger.Logger,
	verifier pow.Verifier,
	repo repository.Repositories,
) *Server {
	return &Server{
		cfg:         cfg,
		logger:      logger,
		verifier:    verifier,
		repo:        repo,
		connections: make(chan net.Conn, cfg.Workers),
	}
}

// Run starts the Server.
func (s *Server) Run(ctx context.Context) (err error) {
	lc := net.ListenConfig{
		KeepAlive: s.cfg.KeepAlive,
	}
	s.listener, err = lc.Listen(ctx, "tcp", s.cfg.Addr)

	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	s.logger.Info(fmt.Sprintf("Server started on port %s", s.listener.Addr().String()))
	go s.handleConnections(ctx) // start workers watching connections queue
	s.serve()

	return nil
}

// Stop graceful server shutdown.
func (s *Server) Stop() {
	err := s.listener.Close()
	if err != nil && !errors.Is(err, net.ErrClosed) {
		s.logger.Error("failed to close listener: ", err.Error())
	}

	s.wg.Wait()
	s.logger.Info("Server stopped")
}

func (s *Server) serve() {
	for {
		conn, err := s.listener.Accept()
		if errors.Is(err, net.ErrClosed) {
			s.logger.Debug("listener closed")
			return
		} else if err != nil {
			s.logger.Error("failed to accept connection: ", err.Error())
			continue
		}

		select {
		case s.connections <- conn:
		case <-time.After(s.cfg.ConnAcceptTimeout):
			s.logger.Warn("The connection queue is full, a server can't handle the incoming connections in time. Skip connection")

			if err := conn.Close(); err != nil {
				s.logger.Error(err)
			}
		}
	}
}
