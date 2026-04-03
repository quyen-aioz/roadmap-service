package main

import (
	"context"
	"fmt"
	"log"
	"roadmap/app/internal/core/serverconfig"
	roadmapmodel "roadmap/app/internal/modules/roadmap/model"
	roadmapgroupmodel "roadmap/app/internal/modules/roadmapgroup/model"
	usermodel "roadmap/app/internal/modules/user/model"
	userrepo "roadmap/app/internal/modules/user/repo"
	userservice "roadmap/app/internal/modules/user/service"
	"roadmap/pkg/jwtx"
	"roadmap/pkg/sqlitex"

	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	if err := db.AutoMigrate(&roadmapmodel.Roadmap{}, &roadmapgroupmodel.RoadmapGroup{}, &usermodel.User{}); err != nil {
		return fmt.Errorf("failed to migrate table: %w", err)
	}

	if err := seed(db); err != nil {
		return fmt.Errorf("failed to seed: %w", err)
	}
	return nil
}

func seed(db *gorm.DB) error {
	adminConf := serverconfig.Get().SeedAdmin
	userSvc := userservice.NewWithRepo(userrepo.New(db))
	if err := userSvc.SeedAdmin(context.Background(), adminConf.Username, adminConf.Password); err != nil {
		return fmt.Errorf("failed to seed admin: %w", err)
	}

	if err := seedGroups(db); err != nil {
		return fmt.Errorf("failed to seed groups: %w", err)
	}

	log.Println("Seeded groups")
	return nil
}

func seedGroups(db *gorm.DB) error {
	groups := []roadmapgroupmodel.RoadmapGroup{
		{ID: roadmapmodel.GroupIDAiozNetwork, Name: "AIOZ Network", SvgURL: "https://content.aioz.network/logo/svg/dark/logo-aioz_network_md.svg", SortOrder: 0},
		{ID: roadmapmodel.GroupIDAiozDepin, Name: "AIOZ DePIN", SvgURL: "https://content.aioz.network/logo/svg/light/logo-aioz_depin_md.svg", SortOrder: 1},
		{ID: roadmapmodel.GroupIDAiozAi, Name: "AIOZ AI", SvgURL: "https://content.aioz.network/logo/svg/light/logo-aioz_ai_md.svg", SortOrder: 2},
		{ID: roadmapmodel.GroupIDAiozStream, Name: "AIOZ Stream", SvgURL: "https://content.aioz.network/logo/svg/light/logo-aioz_stream_md.svg", SortOrder: 3},
		{ID: roadmapmodel.GroupIDAiozStorage, Name: "AIOZ Storage", SvgURL: "https://content.aioz.network/logo/svg/light/logo-aioz_storage_md.svg", SortOrder: 4},
		{ID: roadmapmodel.GroupIDAiozPin, Name: "AIOZ Pin", SvgURL: "https://content.aioz.network/logo/svg/light/logo-aioz_pin_md.svg", SortOrder: 5},
		{ID: roadmapmodel.GroupIDAiozWallet, Name: "AIOZ Wallet", SvgURL: "https://content.aioz.network/logo/svg/light/logo-aioz_wallet_md.svg", SortOrder: 6},
		{ID: roadmapmodel.GroupIDAiozAds, Name: "AIOZ Ads", SvgURL: "https://content.aioz.network/logo/svg/light/logo-aioz_ads_md.svg", SortOrder: 7},
		{ID: roadmapmodel.GroupIDAiozAiAgents, Name: "AIOZ AI Agents", SvgURL: "https://content.aioz.network/logo/svg/light/logo-aioz_agents_md.svg", SortOrder: 8},
		{ID: roadmapmodel.GroupIDAiozBridge, Name: "AIOZ Bridge", SvgURL: "https://content.aioz.network/logo/svg/light/logo-aioz_bridge_md.svg", SortOrder: 9},
		{ID: roadmapmodel.GroupIDAiozDex, Name: "AIOZ Dex", SvgURL: "https://content.aioz.network/logo/svg/light/logo-aioz_dex_md.svg", SortOrder: 10},
		{ID: roadmapmodel.GroupIDAiozExplorer, Name: "AIOZ Explorer", SvgURL: "https://content.aioz.network/logo/svg/light/logo-aioz_explorer_md.svg", SortOrder: 11},
		{ID: roadmapmodel.GroupIDAiozTransfer, Name: "AIOZ Transfer", SvgURL: "https://content.aioz.network/logo/svg/light/logo-aioz_transfer_md.svg", SortOrder: 12},
		{ID: roadmapmodel.GroupIDAiozVault, Name: "AIOZ Vault", SvgURL: "https://content.aioz.network/logo/svg/light/logo-aioz_vault_md.svg", SortOrder: 13},
	}

	return db.Clauses(clause.OnConflict{DoNothing: true}).Create(&groups).Error
}

func initJWT() error {
	conf := serverconfig.Get().JWT

	if err := jwtx.InitJWT(conf.SigningKey); err != nil {
		return fmt.Errorf("failed to init jwt: %w", err)
	}
	return nil
}
