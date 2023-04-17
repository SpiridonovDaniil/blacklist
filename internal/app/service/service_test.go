package service

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	server "blacklist/internal/app/http"
	mock_http "blacklist/internal/app/http/mocks"
	"blacklist/internal/config"
	"blacklist/internal/domain"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var cfg = &config.Config{Auth: config.Auth{Auth: "admin"}}

func TestService_Create(t *testing.T) {
	type mockBehavior func(s *mock_http.Mockservice, expectedError error)
	testTable := []struct {
		name                   string
		inputBody              string
		mockBehavior           mockBehavior
		expectedTestStatusCode int
		expectedError          error
		expectedResponse       string
	}{
		{
			name:      "create HTTP status 201",
			inputBody: `{"name" : "my_name", "phone" : "my_phone", "reason" : "my_reason", "uploader" : "admin"}`,
			mockBehavior: func(s *mock_http.Mockservice, expectedError error) {
				s.EXPECT().Create(gomock.Any(), gomock.AssignableToTypeOf(&domain.Person{})).Return(expectedError)
			},
			expectedTestStatusCode: 201,
			expectedResponse:       "",
		},
		{
			name:                   "create bad request",
			inputBody:              `{"name" : "", "phone" : "my_phone", "reason" : "my_reason", "uploader" : "admin"}`,
			expectedTestStatusCode: 400,
			expectedResponse:       "[createHandler] bad request, name, phone number, reason and adding user fields must be filled in",
		},
		{
			name:                   "create bad request, syntax error",
			inputBody:              `{"name : "my_name", "phone" : "my_phone", "reason" : "my_reason", "uploader" : "admin"}`,
			expectedTestStatusCode: 400,
			expectedResponse:       "[createHandler] failed to parse request, error: invalid character 'm' after object key",
		},
		{
			name:      "create internal server error",
			inputBody: `{"name" : "my_name", "phone" : "my_phone", "reason" : "my_reason", "uploader" : "admin"}`,
			mockBehavior: func(s *mock_http.Mockservice, expectedError error) {
				s.EXPECT().Create(gomock.Any(), gomock.AssignableToTypeOf(&domain.Person{})).Return(expectedError)
			},
			expectedTestStatusCode: 500,
			expectedError:          fmt.Errorf("[create] error adding a user to the blacklist"),
			expectedResponse:       "[createHandler] [create] error adding a user to the blacklist",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			service := mock_http.NewMockservice(c)
			if testCase.name == "create HTTP status 201" {
				testCase.mockBehavior(service, testCase.expectedError)
			}
			if testCase.name == "create internal server error" {
				testCase.mockBehavior(service, testCase.expectedError)
			}

			f := server.NewServer(service, cfg)
			req, err := http.NewRequest("POST", "/", strings.NewReader(testCase.inputBody))
			req.Header.Add("content-Type", "application/json")
			req.Header.Add("Authorization", "admin")
			assert.NoError(t, err)

			resp, err := f.Test(req)
			assert.NoError(t, err)

			body, err := io.ReadAll(resp.Body)
			assert.Equal(t, testCase.expectedTestStatusCode, resp.StatusCode)
			assert.Equal(t, testCase.expectedResponse, string(body))
		})
	}
}

func TestService_Delete(t *testing.T) {
	type mockBehavior func(s *mock_http.Mockservice, expectedError error)
	testTable := []struct {
		name                   string
		inputBody              string
		mockBehavior           mockBehavior
		expectedTestStatusCode int
		expectedError          error
		expectedResponse       string
	}{
		{
			name:      "create HTTP status 200",
			inputBody: `{"id" : 1}`,
			mockBehavior: func(s *mock_http.Mockservice, expectedError error) {
				s.EXPECT().Delete(gomock.Any(), gomock.AssignableToTypeOf(&domain.Id{})).Return(expectedError)
			},
			expectedTestStatusCode: 200,
			expectedResponse:       "",
		},
		{
			name:                   "create bad request, syntax error",
			inputBody:              `{"id" : "1"}`,
			expectedTestStatusCode: 400,
			expectedResponse:       "[deleteHandler] failed to parse request, error: json: cannot unmarshal string into Go struct field Id.id of type int",
		},
		{
			name:      "create internal server error",
			inputBody: `{"id" : 1}`,
			mockBehavior: func(s *mock_http.Mockservice, expectedError error) {
				s.EXPECT().Delete(gomock.Any(), gomock.AssignableToTypeOf(&domain.Id{})).Return(expectedError)
			},
			expectedTestStatusCode: 500,
			expectedError:          fmt.Errorf("[delete] error deleting a person from the blacklist"),
			expectedResponse:       "[deleteHandler] [delete] error deleting a person from the blacklist",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			service := mock_http.NewMockservice(c)
			if testCase.name == "create HTTP status 200" {
				testCase.mockBehavior(service, testCase.expectedError)
			}
			if testCase.name == "create internal server error" {
				testCase.mockBehavior(service, testCase.expectedError)
			}

			f := server.NewServer(service, cfg)
			req, err := http.NewRequest("DELETE", "/", strings.NewReader(testCase.inputBody))
			req.Header.Add("content-Type", "application/json")
			req.Header.Add("Authorization", "admin")
			assert.NoError(t, err)

			resp, err := f.Test(req)
			assert.NoError(t, err)

			body, err := io.ReadAll(resp.Body)
			assert.Equal(t, testCase.expectedTestStatusCode, resp.StatusCode)
			assert.Equal(t, testCase.expectedResponse, string(body))
		})
	}
}

func TestService_Get(t *testing.T) {
	type mockBehavior func(s *mock_http.Mockservice, serviceAnswer []*domain.Person, expectedError error)
	testTable := []struct {
		name                   string
		serviceAnswer          []*domain.Person
		mockBehavior           mockBehavior
		expectedTestStatusCode int
		expectedError          error
		expectedResponse       string
	}{
		{
			name: "create HTTP status 200 name",
			serviceAnswer: []*domain.Person{
				{
					Id:       1,
					Name:     "my_name",
					Phone:    "my_phone1",
					Reason:   "my_reason",
					Time:     "15-04-2023",
					Uploader: "admin",
				},
				{
					Id:       2,
					Name:     "my_name",
					Phone:    "my_phone2",
					Reason:   "another_reason",
					Time:     "16-04-2023",
					Uploader: "admin",
				},
			},
			mockBehavior: func(s *mock_http.Mockservice, serviceAnswer []*domain.Person, expectedError error) {
				s.EXPECT().Get(gomock.Any(), gomock.Any()).Return(serviceAnswer, expectedError)
			},
			expectedTestStatusCode: 200,
			expectedResponse:       `[{"id":1,"phone":"my_phone1","name":"my_name","reason":"my_reason","time":"15-04-2023","uploader":"admin"},{"id":2,"phone":"my_phone2","name":"my_name","reason":"another_reason","time":"16-04-2023","uploader":"admin"}]`,
		},
		{
			name: "create HTTP status 200 phone",
			serviceAnswer: []*domain.Person{
				{
					Id:       1,
					Name:     "my_name1",
					Phone:    "my_phone",
					Reason:   "my_reason",
					Time:     "15-04-2023",
					Uploader: "admin",
				},
				{
					Id:       2,
					Name:     "my_name2",
					Phone:    "my_phone",
					Reason:   "another_reason",
					Time:     "16-04-2023",
					Uploader: "admin",
				},
			},
			mockBehavior: func(s *mock_http.Mockservice, serviceAnswer []*domain.Person, expectedError error) {
				s.EXPECT().Get(gomock.Any(), gomock.Any()).Return(serviceAnswer, expectedError)
			},
			expectedTestStatusCode: 200,
			expectedResponse:       `[{"id":1,"phone":"my_phone","name":"my_name1","reason":"my_reason","time":"15-04-2023","uploader":"admin"},{"id":2,"phone":"my_phone","name":"my_name2","reason":"another_reason","time":"16-04-2023","uploader":"admin"}]`,
		},
		{
			name:                   "create bad request",
			expectedTestStatusCode: 400,
			expectedResponse:       "[getHandler] search parameters are not specified",
		},
		{
			name: "create internal server error",
			mockBehavior: func(s *mock_http.Mockservice, serviceAnswer []*domain.Person, expectedError error) {
				s.EXPECT().Get(gomock.Any(), gomock.AssignableToTypeOf(&domain.Search{})).Return(nil, expectedError)
			},
			expectedTestStatusCode: 500,
			expectedError:          fmt.Errorf("[get] error searching for a person in the blacklist"),
			expectedResponse:       "[getHandler] [get] error searching for a person in the blacklist",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			service := mock_http.NewMockservice(c)
			if testCase.name == "create HTTP status 200 name" || testCase.name == "create HTTP status 200 phone" {
				testCase.mockBehavior(service, testCase.serviceAnswer, testCase.expectedError)
			}
			if testCase.name == "create internal server error" {
				testCase.mockBehavior(service, nil, testCase.expectedError)
			}

			f := server.NewServer(service, cfg)
			var url string
			if testCase.name == "create bad request" {
				url = "/"
			} else {
				url = "/?name=my_name?phone=my_phone"
			}
			req, err := http.NewRequest("GET", url, strings.NewReader(""))
			req.Header.Add("content-Type", "application/json")
			req.Header.Add("Authorization", "admin")
			assert.NoError(t, err)

			resp, err := f.Test(req)
			assert.NoError(t, err)

			body, err := io.ReadAll(resp.Body)
			assert.Equal(t, testCase.expectedTestStatusCode, resp.StatusCode)
			assert.Equal(t, testCase.expectedResponse, string(body))
		})
	}
}
