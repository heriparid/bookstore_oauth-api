package rest

import (
	"net/http"
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

// This library is not support for golang version > 1.12

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "/users/login",
		ReqBody:      `{"email":"email@email.com","password":"password"}`,
		RespHTTPCode: http.StatusInternalServerError,
		RespBody:     "{}",
	})
	repository := restUserRepository{}
	user, err := repository.LoginUser("email@email.com", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {

}
