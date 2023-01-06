package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/walletkita/wallet-core/app"
	"github.com/walletkita/wallet-core/internal/config"
)

type IServer interface {
	Run()
}

type serverContext struct {
	cfg     config.Config
	handler http.Handler
}

func NewServer(cfg config.Config, handler http.Handler) IServer {
	return serverContext{cfg: cfg}
}

func (s serverContext) Run() {
	// Set up a channel to listen to for interrupt signals
	var runChan = make(chan os.Signal, 1)

	// Set up a context to allow for graceful server shutdowns in the event
	// of an OS interrupt (defers the cancel just in case)
	ctx, cancel := context.WithTimeout(
		context.Background(),
		s.cfg.Server.Timeout.Server,
	)
	defer cancel()

	// Define server options
	server := &http.Server{
		Addr:         s.cfg.Server.Host + ":" + s.cfg.Server.Port,
		Handler:      s.handler,
		ReadTimeout:  s.cfg.Server.Timeout.Read * time.Second,
		WriteTimeout: s.cfg.Server.Timeout.Write * time.Second,
		IdleTimeout:  s.cfg.Server.Timeout.Idle * time.Second,
	}

	// Handle ctrl+c/ctrl+x interrupt
	signal.Notify(runChan, os.Interrupt, syscall.SIGTERM)

	// Alert the user that the server is starting
	app.LOG.Debugf("Server is starting on %s\n", server.Addr)

	// Run the server on a new goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				// Normal interrupt operation, ignore
			} else {
				app.LOG.Errorf("Server failed to start due to err: %v", err)
			}
		}
	}()

	// Block on this channel listeninf for those previously defined syscalls assign
	// to variable so we can let the user know why the server is shutting down
	interrupt := <-runChan

	// If we get one of the pre-prescribed syscalls, gracefully terminate the server
	// while alerting the user
	app.LOG.Warnf("Server is shutting down due to %+v\n", interrupt)

	if err := server.Shutdown(ctx); err != nil {
		app.LOG.Fatalf("Server was unable to gracefully shutdown due to err: %+v", err)
	}
}
