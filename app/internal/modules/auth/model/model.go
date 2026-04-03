package authmodel

type (
	LoginReq struct {
		Username string
		Password string
	}

	LoginResp struct {
		AccessToken string
	}
)

type (
	RegisterReq struct {
		Username string
		Password string
	}

	RegisterResp struct {
		AccessToken string
	}
)

type (
	ChangePasswordReq struct {
		OldPassword string
		NewPassword string
	}

	ChangePasswordResp struct {
		AccessToken string
	}
)
