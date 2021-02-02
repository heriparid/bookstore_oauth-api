package access_token

import (
	"strings"
	"time"

	"github.com/heriparid/oauth-api/src/utils/errors"
)

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

// Validate value
func (at *AccessToken) Validate() *errors.RestErr {
	at.AcccessToken = strings.TrimSpace(at.AcccessToken)
	if at.AcccessToken == "" {
		return errors.NewBadRequestError("invalid access token")
	}

	if at.UserID <= 0 {
		return errors.NewBadRequestError("invalid user ID")
	}

	if at.ClientID <= 0 {
		return errors.NewBadRequestError("invalid client ID")
	}

	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiry time")
	}

	return nil
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
