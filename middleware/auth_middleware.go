package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"projectsistemkasir/helper"
	"projectsistemkasir/model/web"
	"strings"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// if "PINJAM_DULU_SERATUS" == request.Header.Get("Authorization") {
		
	// 	middleware.Handler.ServeHTTP(writer, request) //Diteruskan ke handler selanjutnya

	// } else {
	// 	// error
	// 	writer.Header().Set("Content-Type", "application/json")
	// 	writer.WriteHeader(http.StatusUnauthorized)

	// 	webResponse := web.WebResponse {
	// 		Code: http.StatusUnauthorized,
	// 		Status: "UNAUTHORIZED",
	// 	}

	// 	helper.WriteToResponseBody(writer, webResponse)
		
	// }


	writer.Header().Set("Content-Type", "application/json")
	token := request.Header.Get("Authorization")

	if token == "" {
		writer.WriteHeader(http.StatusUnauthorized)
		webResponse := web.WebResponse {
			Code: http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}
		json.NewEncoder(writer).Encode(webResponse)
		return
	}

	token = strings.Split(token, " ")[1]
	claims, err := helper.DecodeToken(token)

	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		webResponse := web.WebResponse {
			Code: http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}
		json.NewEncoder(writer).Encode(webResponse)
		return
	}

	ctx := context.WithValue(request.Context(), "userInfo", claims)
	request = request.WithContext(ctx)
	middleware.Handler.ServeHTTP(writer, request.WithContext(ctx))

}

//cara implementasi middleware di golang (middleware untuk jwt auth)