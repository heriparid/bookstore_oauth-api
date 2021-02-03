package rest

import (
	"encoding/json"
	"time"

	"github.com/heriparid/oauth-api/src/domain/users"
	"github.com/heriparid/oauth-api/src/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "localhost:8080",
		Timeout: 100 * time.Microsecond,
	}
)

//RestUsersRepository interface
type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type restUserRepository struct {
}

// NewRepository instance
func NewRepository() RestUsersRepository {
	return &restUserRepository{}
}

func (r *restUserRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {

	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	response := usersRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("invalid rest client response when trying to hit login API")
	}

	if response.StatusCode > 299 {
		var restErr errors.RestErr

		if err := json.Unmarshal(response.Bytes(), &restErr); err != nil {
			return nil, errors.NewInternalServerError("invalid error interface when trying to login user")
		}

		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("invalid user interface when trying to login user")
	}

	return &user, nil
}
