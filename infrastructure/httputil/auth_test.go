package httputil

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-chi/chi"

	testHelper "github.com/nekochans/kimono-app-api/test"
)

var region string
var userPoolId string
var userPoolClientId string
var accessToken *string

func TestMain(m *testing.M) {
	region = os.Getenv("REGION")
	userPoolId = os.Getenv("USER_POOL_ID")
	userPoolClientId = os.Getenv("USER_POOL_WEB_CLIENT_ID")
	email := os.Getenv("TEST_EMAIL")
	password := os.Getenv("TEST_PASSWORD")

	helper := testHelper.CognitoAuthHelper{
		Region:           region,
		UserPoolID:       userPoolId,
		UserPoolClientID: userPoolClientId,
	}

	_, err := helper.SignUp(email, password)
	if err != nil {
		fmt.Printf("fail to signUp: %s\n", err)
		os.Exit(1)
	}

	authOutput, err := helper.SignIn(email, password)
	if err != nil {
		fmt.Printf("fail to signIp: %s\n", err)
		os.Exit(1)
	}

	accessToken = authOutput.AuthenticationResult.AccessToken

	status := m.Run()

	err = helper.DeleteUser(*accessToken)
	if err != nil {
		fmt.Printf("fail to delete user: %s\n", err)
		os.Exit(1)
	}

	// TODO Go 1.15では os.Exit 不要
	os.Exit(status)
}

func TestAuth(t *testing.T) {
	tests := []struct {
		name               string
		request            func() *http.Request
		userPoolId         string
		expectedResponse   string
		expectedStatusCode int
	}{
		{
			"OK response when including the Authorization header",
			func() *http.Request {
				req, _ := http.NewRequestWithContext(context.TODO(), http.MethodGet, "/", nil)
				req.Header.Add("Authorization", *accessToken)
				req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64)")
				return req
			},
			userPoolId,
			"Hello, HTTPサーバ",
			http.StatusOK,
		},
		{
			"Error response when the Authorization header is not included",
			func() *http.Request {
				req, _ := http.NewRequestWithContext(context.TODO(), http.MethodGet, "/", nil)
				req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64)")
				return req
			},
			userPoolId,
			"Unauthorized",
			http.StatusUnauthorized,
		},
		{
			"Error response when failed to fetch JWK",
			func() *http.Request {
				req, _ := http.NewRequestWithContext(context.TODO(), http.MethodGet, "/", nil)
				req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64)")
				return req
			},
			"invalid-user-poop-id",
			"Internal Server Error",
			http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := chi.NewRouter()

			r.Use(Auth(region, test.userPoolId))
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, "Hello, HTTPサーバ")
			})
			r.ServeHTTP(w, test.request())

			if w.Body.String() != test.expectedResponse {
				t.Error("response Body was not the expected value")
				t.Error("\nActually: ", w.Body.String(), "\nExpected: ", test.expectedResponse)
			}
			if w.Code != test.expectedStatusCode {
				t.Error("status code was not the expected value")
				t.Error("\nActually: ", w.Code, "\nExpected: ", test.expectedStatusCode)
			}
		})
	}
}
