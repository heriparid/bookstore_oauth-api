package db

import (
	"github.com/gocql/gocql"
	"github.com/heriparid/oauth-api/src/client/cassandra"
	"github.com/heriparid/oauth-api/src/domain/access_token"
	"github.com/heriparid/oauth-api/src/utils/errors"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES(?,?,?,?);"
	queryUpdateExpiryTime  = "UPDATE access_tokens SET expires=? WHERE access_token=?"
)

// DbRepository interface
type DbRepository interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpiryTime(access_token.AccessToken) *errors.RestErr
}
type dbRepository struct {
}

// NewRepository new instance
func NewRepository() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetByID(accessTokenID string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer session.Close()

	var result access_token.AccessToken
	if err := session.Query(queryGetAccessToken, accessTokenID).Scan(
		&result.AcccessToken,
		&result.UserID,
		&result.ClientID,
		&result.Expires,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundRequestError("no access token found with given id")
		}
	}

	return &result, nil
}

func (r *dbRepository) Create(at access_token.AccessToken) *errors.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer session.Close()

	if err := session.Query(queryCreateAccessToken,
		at.AcccessToken,
		at.UserID,
		at.ClientID,
		at.Expires,
	).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (r *dbRepository) UpdateExpiryTime(at access_token.AccessToken) *errors.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer session.Close()

	if err := session.Query(queryUpdateExpiryTime,
		at.Expires,
		at.AcccessToken,
	).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}
