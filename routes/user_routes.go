package routes

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	userproto "github.com/vivaldy22/mekar-regis-client/proto"
	"github.com/vivaldy22/mekar-regis-client/tools/respJson"
	"github.com/vivaldy22/mekar-regis-client/tools/vError"
	"github.com/vivaldy22/mekar-regis-client/tools/varMux"
	"net/http"
)

type userRoute struct {
	service userproto.UserCRUDClient
}

func NewUserRoute(service userproto.UserCRUDClient, r *mux.Router) {
	handler := &userRoute{service}

	prefix := r.PathPrefix("/users").Subrouter()
	prefix.HandleFunc("", handler.getAll).Methods(http.MethodGet)
}

func (u *userRoute) getAll(w http.ResponseWriter, r *http.Request) {
	data, err := u.service.GetAll(context.Background(), new(userproto.Empty))

	if err != nil {
		vError.WriteError("Get All Users failed!", http.StatusBadRequest, err, w)
	} else {
		respJson.WriteJSON(data.List, w)
	}
}

func (u *userRoute) getByID(w http.ResponseWriter, r *http.Request) {
	id := varMux.GetVarsMux("id", r)

	data, err := u.service.GetByID(context.Background(), &userproto.ID{
		Id: id,
	})

	if err != nil {
		vError.WriteError("Get User by ID failed", http.StatusBadRequest, err, w)
	} else {
		respJson.WriteJSON(data, w)
	}
}

func (u *userRoute) create(w http.ResponseWriter, r *http.Request) {
	var user *userproto.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		vError.WriteError("Decoding json failed!", http.StatusExpectationFailed, err, w)
	} else {
		created, err := u.service.Create(context.Background(), user)

		if err != nil {
			vError.WriteError("Create User Failed!", http.StatusBadRequest, err, w)
		} else {
			data, err := u.service.GetByID(context.Background(), &userproto.ID{
				Id: created.UserId,
			})

			if err != nil {
				vError.WriteError("Get User by ID failed", http.StatusBadRequest, err, w)
			} else {
				respJson.WriteJSON(data, w)
			}
		}
	}
}

func (u *userRoute) update(w http.ResponseWriter, r *http.Request) {
	var user *userproto.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		vError.WriteError("Decoding json failed", http.StatusExpectationFailed, err, w)
	} else {
		id := varMux.GetVarsMux("id", r)

		_, err := u.service.Update(context.Background(), &userproto.UserUpdateRequest{
			Id:   id,
			User: user,
		})

		if err != nil {
			vError.WriteError("Updating data failed!", http.StatusBadRequest, err, w)
		} else {
			data, err := u.service.GetByID(context.Background(), &userproto.ID{
				Id: id,
			})

			if err != nil {
				vError.WriteError("Get User by ID failed!", http.StatusBadRequest, err, w)
			} else {
				respJson.WriteJSON(data, w)
			}
		}
	}
}

func (u *userRoute) delete(w http.ResponseWriter, r *http.Request) {
	id := varMux.GetVarsMux("id", r)

	userID := &userproto.ID{
		Id: id,
	}
	data, err := u.service.GetByID(context.Background(), userID)

	if err != nil {
		vError.WriteError("Get User By ID failed!", http.StatusBadRequest, err, w)
	} else {
		_, err := u.service.Delete(context.Background(), userID)

		if err != nil {
			vError.WriteError("Delete User failed!", http.StatusBadRequest, err, w)
		} else {
			respJson.WriteJSON(data, w)
		}
	}
}