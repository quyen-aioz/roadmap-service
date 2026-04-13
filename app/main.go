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
	warmup           = 30 * time.Second
)

func main() {
	must(setDefaultTimezoneUTC(), "set default timezone UTC")

	serverconfig.Init()
	conf := serverconfig.Get()
	httpx.DebugMsgEnabled = serverconfig.IsNonPROD()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	timeoutContext, cancel := context.WithTimeout(ctx, warmup)
	defer cancel()

	must(init3rdParties(timeoutContext), "init 3rd parties")
	cronRunner.Start(ctx)

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

		cronRunner.Stop()

		return nil
	})

	if err := errg.Wait(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("server shutdown in error: %w", err)
	}

	log.Println("server shutdown gracefully")
}

func setDefaultTimezoneUTC() error {
	time.Local = time.UTC
	return nil
}

func must(err error, errMsg string) {
	if err == nil {
		return
	}
	log.Fatalf("%s: %v", errMsg, err)
}
