package middleware

import (
	"net/http"

	"github.com/rifqimuhammadaziz/go-restful-api/helper"
	"github.com/rifqimuhammadaziz/go-restful-api/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if "SECRETKEY" == r.Header.Get("X-API-Key") {
		// ok
		middleware.Handler.ServeHTTP(rw, r)
	} else {
		// error
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusUnauthorized) // set status code(500)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(rw, webResponse)
	}
}
