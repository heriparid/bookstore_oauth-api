package db

import (
	"github.com/heriparid/oauth-api/src/domain/access_token"
	"github.com/heriparid/oauth-api/src/utils/errors"
)

// DbRepository interface
type DbRepository interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
}
type dbRepository struct {
}

// NewRepository new instance
func NewRepository() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetByID(accessTokenID string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("database connection is not implemented yet")
}
