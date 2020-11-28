package httputil

import (
	"fmt"
	"net/http"

	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

func Auth(region, userPoolId, aud string) func(next http.Handler) http.Handler {
	iss := fmt.Sprintf(
		"https://cognito-idp.%v.amazonaws.com/%v",
		region,
		userPoolId,
	)
	jwkURL := fmt.Sprintf("%v/.well-known/jwks.json", iss)

	keySet, err := jwk.FetchHTTP(jwkURL)
	if err != nil {
		fmt.Printf("failed to fetch JWK: %s\n", err)
	}

	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte("Internal Server Error"))
				return
			}

			header := r.Header
			tokenString, ok := header["Authorization"]
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				_, _ = w.Write([]byte("Unauthorized"))
				return
			}

			_, err := validateToken(tokenString[0], iss, aud, keySet)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				_, _ = w.Write([]byte("Unauthorized"))
				return
			}

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

func validateToken(tokenStr, iss, aud string, keySet *jwk.Set) (jwt.Token, error) {
	token, err := jwt.ParseString(
		tokenStr,
		jwt.WithKeySet(keySet),
	)
	if err != nil {
		fmt.Printf("failed to parse JWT token: %s\n", err)
		return nil, err
	}

	err = jwt.Verify(
		token,
		jwt.WithIssuer(iss),
		jwt.WithAudience(aud),
		jwt.WithClaimValue("token_use", "id"),
	)

	if err != nil {
		fmt.Printf("failed to verify JWT token: %s\n", err)
		return nil, err
	}
	return token, nil
}
