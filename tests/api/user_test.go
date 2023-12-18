package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/goravel/framework/facades"
	"github.com/stretchr/testify/suite"

	"goravel/app/models"
	"goravel/tests"
)

type UserTestSuite struct {
	suite.Suite
	tests.TestCase
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

// SetupTest will run before each test in the suite.
func (s *UserTestSuite) SetupTest() {
	// s.RefreshDatabase()
}

// TearDownTest will run after each test in the suite.
func (s *UserTestSuite) TearDownTest() {
}

type ApiResponse struct {
	Message string        `json:"message"`
	Data    []models.User `json:"data"`
}

func (s *UserTestSuite) TestIndex() {
	// TODO
	var user models.User
	err := facades.Orm().Factory().Create(&user)
	s.NoError(err)

	resp, err := http.Get("http://localhost:3000/users")
	s.NoError(err)

	s.Equal(200, resp.StatusCode)

	var apiResponse ApiResponse
	var users []models.User
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	users = apiResponse.Data
	resp.Body.Close()
	s.NoError(err)

	s.Equal(users[0].ID, user.ID)
	s.Equal(users[0].Full_Name, user.Full_Name)
	s.Equal(users[0].Email, user.Email)
	s.Equal(users[0].Role, user.Role)
	s.Equal(users[0].Is_Active, user.Is_Active)
}

func (s *UserTestSuite) TestShow() {
	// TODO
	var user models.User
	err := facades.Orm().Factory().Create(&user)
	s.NoError(err)

	id := fmt.Sprintf("%d", user.ID)
	resp, err := http.Get("http://localhost:3000/users/" + id)
	s.NoError(err)

	s.Equal(200, resp.StatusCode)

	var apiResponse ApiResponse
	var users []models.User
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	users = apiResponse.Data
	resp.Body.Close()
	s.NoError(err)

	s.Equal(users[0].ID, user.ID)
	s.Equal(users[0].Full_Name, user.Full_Name)
	s.Equal(users[0].Email, user.Email)
	s.Equal(users[0].Role, user.Role)
	s.Equal(users[0].Is_Active, user.Is_Active)
}
