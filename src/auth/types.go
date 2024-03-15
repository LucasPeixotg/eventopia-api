package auth

type tokenResponse struct {
	Token string `json:"token"`
	Name  string `json:"name"`
}
