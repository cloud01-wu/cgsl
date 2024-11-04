package model

type UserSession struct {
	Username   string `json:"username"`
	CryptoKey  string `json:"cryptoKey"`
	SessionID  string `json:"sessionId"`
	IssuedTime int64  `json:"issuedTime"`
}
