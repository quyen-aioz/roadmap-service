package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"roadmap/app/internal/core/serverconfig"
	"roadmap/app/routes"
	"roadmap/pkg/httpx"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

const (
	gracefulShutdown = 15 * time.Second
)

func main() {
	httpx.DebugMsgEnabled = true

	serverconfig.Init()
	conf := serverconfig.Get()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	e := routes.New()

	errg, runningContext := errgroup.WithContext(ctx)

	errg.Go(func() error {
		err := e.Start(fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port))
		if !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("failed to start server: %w", err)
		}

		return nil
	})

	errg.Go(func() error {
		<-runningContext.Done()

		ctx, cancel := context.WithTimeout(context.Background(), gracefulShutdown)
		defer cancel()

		if err := e.Shutdown(ctx); err != nil {
			return fmt.Errorf("shutdown http server: %w", err)
		}

		return nil
	})

	if err := errg.Wait(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("server shutdown in error: %w", err)
	}

	log.Println("server shutdown gracefully")
}
