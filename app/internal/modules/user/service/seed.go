package userservice

import (
	"context"
	"log"
	usermodel "roadmap/app/internal/modules/user/model"
)

func (svc *Service) SeedAdmin(ctx context.Context, username, password string) error {
	if username == "" || password == "" {
		return nil // quyen@note: skip seed if username or password is empty
	}

	_, err := svc.GetUserByUsername(ctx, username)
	if err == nil {
		return nil // quyen@note: skip seed if user already exists
	}

	_, err = svc.CreateUser(ctx, usermodel.CreateUserReq{
		Username: username,
		Password: password,
	})
	if err != nil {
		return err
	}

	log.Printf("Seeded admin user: %s\n", username)
	return nil
}
