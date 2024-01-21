package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"projectsistemkasir/helper"
	"projectsistemkasir/model/web"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func AuthJWTMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		//logika jwt
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

		ctx := context.WithValue(request.Context(), "userId", claims)
		request = request.WithContext(ctx)

		next(writer,request,params)
	}
}