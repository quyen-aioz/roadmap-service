package main

import (
	"context"
	"fmt"
	"roadmap/app/internal/core/serverconfig"
	roadmapmodel "roadmap/app/internal/modules/roadmap/model"
	usermodel "roadmap/app/internal/modules/user/model"
	userrepo "roadmap/app/internal/modules/user/repo"
	userservice "roadmap/app/internal/modules/user/service"
	"roadmap/pkg/jwtx"
	"roadmap/pkg/sqlitex"

	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
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

	db, err := sqlitex.InitDB(dbPath)
	if err != nil {
		return fmt.Errorf("failed to init db: %w", err)
	}

	if err := db.AutoMigrate(&roadmapmodel.Roadmap{}, &usermodel.User{}); err != nil {
		return fmt.Errorf("failed to migrate table: %w", err)
	}

	if err := seedAdmin(db); err != nil {
		return fmt.Errorf("failed to seed admin: %w", err)
	}
	return nil
}

func seedAdmin(db *gorm.DB) error {
	adminConf := serverconfig.Get().SeedAdmin
	userSvc := userservice.NewWithRepo(userrepo.New(db))
	if err := userSvc.SeedAdmin(context.Background(), adminConf.Username, adminConf.Password); err != nil {
		return fmt.Errorf("failed to seed admin: %w", err)
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
