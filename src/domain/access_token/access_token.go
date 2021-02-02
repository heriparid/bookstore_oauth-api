package access_token

import "time"

const (
	expirationTime = 24
)

// AccessToken domain
type AccessToken struct {
	AcccessToken string `json:"access_token"`
	UserID       int64  `json:"user_id"`
	ClientID     int64  `json:"client_id"`
	Expires      int64  `json:"expires"`
}

// GetNewAccessToken for creating new instance
func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

// IsExpired func
func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
