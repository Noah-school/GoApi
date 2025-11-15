package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/noah-school/goapi/api"
	"github.com/noah-school/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var errUnAuthorized = errors.New("invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		token := r.Header.Get("token")

		if token == "" {
			auth := r.Header.Get("Authorization")
			if auth != "" {
				low := strings.ToLower(auth)
				if strings.HasPrefix(low, "bearer ") {
					token = strings.TrimSpace(auth[7:])
				} else {
					token = strings.TrimSpace(auth)
				}
			}
		}

		if username == "" || token == "" {
			log.Error(errUnAuthorized)
			api.RequestErrorHandler(w, errUnAuthorized)
			return
		}

		database, err := tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		loginDetails := (*database).GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(errUnAuthorized)
			api.RequestErrorHandler(w, errUnAuthorized)
			return
		}
		next.ServeHTTP(w, r)
	})

}
