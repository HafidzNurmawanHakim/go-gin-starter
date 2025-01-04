package middlewares

import (
	"context"
	"gin-template/lib/schema"
	"gin-template/lib/utils"
	"net/http"
)

var jwtKey = []byte("my_secret_key")

func WithJwtAuth(secret string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("access_token")

			if err != nil {
				utils.JSON(w, 401, schema.ErrorResponse{Message: "Missing or invalid token!"})
				return
			}

			access_token := cookie.Value

			if authorize, err := utils.IsAuthorized(access_token, string(jwtKey)); err != nil || !authorize {
				utils.JSON(w,401, schema.ErrorResponse{Message: err.Error()})

			}

			userId, err :=  utils.ExtraxtIdFromToken(access_token, string(jwtKey))

			if err != nil {
				utils.JSON(w, http.StatusUnauthorized, schema.ErrUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "userId", userId)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)

		})
	}
}
