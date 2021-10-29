package requests

type InsertAccessTokenRequest struct {
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}
