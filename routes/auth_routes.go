package routes

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	userproto "github.com/vivaldy22/mekar-regis-client/proto"
	"github.com/vivaldy22/mekar-regis-client/tools/respJson"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type authRoute struct {
	service userproto.AuthRPCClient
}

func NewAuthRoute(service userproto.AuthRPCClient, r *mux.Router) {
	handler := &authRoute{service}

	prefAuth := r.PathPrefix("/auth").Subrouter()
	prefAuth.HandleFunc("", handler.login).Methods(http.MethodPost)
}

func (a *authRoute) login(w http.ResponseWriter, r *http.Request) {
	var credentials *userproto.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&credentials)

	if err != nil {
		respJson.WriteJSON(false, http.StatusBadRequest, "Decode JSON failed", nil, err, w)
	} else {
		pass, err := a.service.GetPassword(context.Background(), credentials)

		if err != nil {
			respJson.WriteJSON(false, http.StatusBadRequest, "Username not found", nil, nil, w)
		} else {
			err = bcrypt.CompareHashAndPassword([]byte(pass.HashedPassword), []byte(credentials.Password))

			if err != nil {
				respJson.WriteJSON(false, http.StatusBadRequest, "Username and password don't match", nil, nil, w)
			} else {
				token, err := a.service.GenerateToken(context.Background(), &userproto.LoginRequest{
					Username: credentials.Username,
					Password: pass.HashedPassword,
				})

				if err != nil {
					respJson.WriteJSON(false, http.StatusExpectationFailed, "Generate Token failed", nil, err, w)
				} else {
					respJson.WriteJSON(true, http.StatusOK, "Login success", &userproto.LoginResponse{
						Token: token.Token,
					}, nil, w)
				}
			}
		}
	}
}