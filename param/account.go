package param

type AddAccountRequest struct {
	AppVersion string `json:"app_version"`
	Email      string `json:"email"`
}

type AddAccountResponse struct {
	AccountID uint          `json:"account_id"`
	Tokens    AccountTokens `json:"tokens"`
}

type AccountTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
