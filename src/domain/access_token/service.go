package access_token

import "github.com/heriparid/oauth-api/src/utils/errors"

// Repository service
type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpiryTime(AccessToken) *errors.RestErr
}

// Service interface
type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpiryTime(AccessToken) *errors.RestErr
}

type service struct {
	repository Repository
}

// NewService create new service instance
func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

// GetById to get AccessToken by ID
func (s *service) GetByID(accessTokenID string) (*AccessToken, *errors.RestErr) {
	at, err := s.repository.GetByID(accessTokenID)
	if err != nil {
		return nil, err
	}
	return at, nil
}

func (s *service) Create(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}

	return s.repository.Create(at)
}

func (s *service) UpdateExpiryTime(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}

	return s.repository.UpdateExpiryTime(at)
}
