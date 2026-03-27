package main

import (
	"context"
	"fmt"
	"roadmap/app/internal/core/serverconfig"
	"roadmap/pkg/sqlitex"

	"golang.org/x/sync/errgroup"
)

func init3rdParties(ctx context.Context) error {
	// quyen@note: overkill since only 1 goroutine -> setup for scaling
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return initSqlite(ctx)
	})

	return g.Wait()
}

func initSqlite(_ context.Context) error {
	conf := serverconfig.Get().SQLite
	dbPath := fmt.Sprintf("%s/%s", conf.Directory, conf.DatabaseName)

	_, err := sqlitex.InitDB(dbPath)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	return nil
}
