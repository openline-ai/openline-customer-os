package model

import "time"

type OAuthToken struct {
	AccessToken       string    `json:"accessToken"`
	RefreshToken      string    `json:"refreshToken"`
	ExpiresAt         time.Time `json:"expiresAt"`
	Scope             string    `json:"scope"`
	ProviderAccountId string    `json:"providerAccountId"`
}

type SignInRequest struct {
	Email      string     `json:"email"`
	Provider   string     `json:"provider"`
	OAuthToken OAuthToken `json:"oAuthToken"`
}
