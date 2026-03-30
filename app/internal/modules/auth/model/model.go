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
