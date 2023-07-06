package handler

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/sharifsharifzoda/project-management-system/models"
	mock_service "github.com/sharifsharifzoda/project-management-system/pkg/service/mocks"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, user models.User)

	testTable := []struct {
		name                 string
		inputBody            string
		inputUser            models.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"firstname":"Test","lastname":"Testov","email":"sharifov300@gmail.com","password":"pass"}`,
			inputUser: models.User{
				Firstname: "Test",
				Lastname:  "Testov",
				Email:     "sharifov300@gmail.com",
				Password:  "pass",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user models.User) {
				s.EXPECT().ValidateUser(user).Return(nil).AnyTimes()
				s.EXPECT().IsEmailUsed(user.Email).Return(false).AnyTimes()
				s.EXPECT().CreateUser(&user).Return(1, nil).AnyTimes()
			},
			expectedStatusCode:   201,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:      "Empty fields",
			inputBody: `{"name":"","email":"","password":""}`,
			inputUser: models.User{},
			mockBehavior: func(s *mock_service.MockAuthorization, user models.User) {
				s.EXPECT().ValidateUser(user).Return(fmt.Errorf("forbidden")).AnyTimes()
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"validate"}`,
		},
		{
			name:      "Used Email",
			inputBody: `{"firstname":"sharif","lastname":"sharifzoda","email":"sharif@gmail.com","password":"qwerty"}`,
			inputUser: models.User{
				Firstname: "sharif",
				Lastname:  "sharifzoda",
				Email:     "sharif@gmail.com",
				Password:  "qwerty",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user models.User) {
				s.EXPECT().ValidateUser(user).Return(nil).AnyTimes()
				s.EXPECT().IsEmailUsed(user.Email).Return(true).AnyTimes()
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"email is already created"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.inputUser)

			//services := &service.Service{Auth: auth}
			handler := NewHandler(auth, nil, nil, nil)
			//handler := Handler{services.Auth, services.Todo}

			gin.SetMode(gin.ReleaseMode)
			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up",
				bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedResponseBody)
		})
	}
}

func TestHandler_signIn(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, user models.User)

	testTable := []struct {
		name                 string
		inputBody            string
		inputUser            models.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"email":"sharifov300@gmail.com","password":"pass"}`,
			inputUser: models.User{
				Email:    "sharifov300@gmail.com",
				Password: "pass",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user models.User) {
				s.EXPECT().GenerateToken(user).Return("token", nil).AnyTimes()
				s.EXPECT().CheckUser(user).Return(user, nil).AnyTimes()
				s.EXPECT().ValidateUser(user).Return(nil).AnyTimes()
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"msg":"signed in"}`,
		},
		{
			name:      "Empty fields",
			inputBody: `{"email":"","password":""}`,
			inputUser: models.User{
				Email:    "",
				Password: "",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user models.User) {
				s.EXPECT().ValidateUser(user).Return(fmt.Errorf("forbidden")).AnyTimes()
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"invalid JSON provided"}`,
		},
		{
			name:      "Invalid fields",
			inputBody: `{"email":"sharifov300@gmail.com","password":"pass"}`,
			inputUser: models.User{
				Email:    "sharifov300@gmail.com",
				Password: "pass",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user models.User) {
				s.EXPECT().CheckUser(user).Return(user, errors.New("invalid email or password")).AnyTimes()
				s.EXPECT().ValidateUser(user).Return(nil).AnyTimes()
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"invalid email or password"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.inputUser)

			handler := NewHandler(auth, nil, nil, nil)

			gin.SetMode(gin.ReleaseMode)
			r := gin.New()
			r.POST("/sign-in", handler.signIn)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-in", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedResponseBody)
		})
	}
}
