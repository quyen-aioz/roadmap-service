package main

import (
	"context"
	"fmt"
	"roadmap/app/internal/core/serverconfig"
	"roadmap/pkg/jwtx"
	"roadmap/pkg/sqlitex"

	"golang.org/x/sync/errgroup"
)

func init3rdParties(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return initSqlite(ctx)
	})
	g.Go(func() error {
		return initJWT()
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
func initJWT() error {
	conf := serverconfig.Get().JWT

	if err := jwtx.InitJWT(conf.SigningKey); err != nil {
		return fmt.Errorf("failed to init jwt: %w", err)
	}
	return nil
}
