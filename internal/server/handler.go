package server

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/OlegBabakov/pow-server/utils"
)

func (s *Server) handleConnections(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case conn := <-s.connections:
			go func() { // run new worker (amount of workers is limited by config)
				_ = conn.SetDeadline(time.Now().Add(s.cfg.ConnDeadline))

				if err := s.handleConnection(conn); err != nil {
					s.logger.Error("connection handle error: ", err.Error())
				}
			}()
		}
	}
}

func (s *Server) handleConnection(conn net.Conn) error {
	s.wg.Add(1)

	defer func() {
		if err := conn.Close(); err != nil {
			s.logger.Error("connection closing error: ", err)
		}

		s.wg.Done()
	}()

	// receive challenge request
	if _, err := utils.ReadMessage(conn); err != nil {
		return fmt.Errorf("read message err: %w", err)
	}

	// send challenge
	challenge := s.verifier.Challenge()
	if err := utils.WriteMessage(conn, challenge); err != nil {
		return fmt.Errorf("send challenge err: %w", err)
	}

	// receive solution
	solution, err := utils.ReadMessage(conn)
	if err != nil {
		return fmt.Errorf("receive proof err: %w", err)
	}

	// verify solution
	if err = s.verifier.Verify(challenge, solution); err != nil {
		return fmt.Errorf("invalid verify: %w", err)
	}

	// send result
	quote, err := s.repo.Quotes.GetQuote()
	if err != nil {
		return fmt.Errorf("get quote err: %w", err)
	}

	if err = utils.WriteMessage(conn, []byte(quote)); err != nil {
		return fmt.Errorf("send quote err: %w", err)
	}

	return nil
}
