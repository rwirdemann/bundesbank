package user

import (
	"testing"

	"fmt"

	"bitbucket.org/rwirdemann/go-workshop/shop/common"
	"bitbucket.org/rwirdemann/go-workshop/shop/user/util"
	"errors"
)

var userService UserService

type MockUserRepository struct{}

func (r MockUserRepository) FindUser(username string) (User, error) {
	if username == "ralf" {
		return User{Username: "Ralf", Password: "test1234"}, nil
	}

	return User{}, errors.New("user not found")
}

func init() {
	userService = UserService{Repository: MockUserRepository{}}
}

func TestSuccessfullAuthentication(t *testing.T) {
	tokenString, ok := Authenticate("ralf", "test1234", userService)
	fmt.Printf("Token: %s", tokenString)
	util.AssertTrue(t, ok)
	username, err := common.ValidateToken(tokenString)
	util.AssertNil(t, err)
	util.AssertEquals(t, "ralf", username)
}

func TestFailedAuthentication(t *testing.T) {
	tokenString, ok := Authenticate("ralf", "test234", userService)
	util.AssertFalse(t, ok)
	util.AssertEquals(t, "", tokenString)
}

func TestTokenManupulation(t *testing.T) {
	tokenString, ok := Authenticate("ralf", "test1234", userService)
	util.AssertTrue(t, ok)
	_, err := ValidateToken(tokenString + "x")
	util.AssertNotNil(t, err)
}
