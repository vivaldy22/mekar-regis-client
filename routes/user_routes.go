package routes

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	userproto "github.com/vivaldy22/mekar-regis-client/proto"
	"github.com/vivaldy22/mekar-regis-client/tools/respJson"
	"github.com/vivaldy22/mekar-regis-client/tools/validation"
	"github.com/vivaldy22/mekar-regis-client/tools/varMux"
	"net/http"
)

type userRoute struct {
	service userproto.UserCRUDClient
}

func NewUserRoute(service userproto.UserCRUDClient, r *mux.Router) {
	handler := &userRoute{service}

	prefUser := r.PathPrefix("/users").Subrouter()
	prefUser.HandleFunc("", handler.getAll).Methods(http.MethodGet)
	prefUser.HandleFunc("/{id}", handler.getByID).Methods(http.MethodGet)
	prefUser.HandleFunc("", handler.create).Methods(http.MethodPost)
	prefUser.HandleFunc("/{id}", handler.update).Methods(http.MethodPut)
	prefUser.HandleFunc("/{id}", handler.delete).Methods(http.MethodDelete)
}

func (u *userRoute) getAll(w http.ResponseWriter, r *http.Request) {
	data, err := u.service.GetAll(context.Background(), new(userproto.Empty))

	if err != nil {
		respJson.WriteJSON(false, http.StatusBadRequest, "Get All Users failed", nil, err, w)
		//vError.WriteError("Get All Users failed!", http.StatusBadRequest, err, w)
	} else {
		respJson.WriteJSON(true, http.StatusOK, "Data found", data.List, nil, w)
	}
}

func (u *userRoute) getByID(w http.ResponseWriter, r *http.Request) {
	id := varMux.GetVarsMux("id", r)

	data, err := u.service.GetByID(context.Background(), &userproto.ID{
		Id: id,
	})

	if err != nil {
		respJson.WriteJSON(false, http.StatusBadRequest, "Get User by ID failed", nil, err, w)
		//vError.WriteError("Get User by ID failed", http.StatusBadRequest, err, w)
	} else {
		respJson.WriteJSON(true, http.StatusOK, "Data found", data, nil, w)
	}
}

func (u *userRoute) create(w http.ResponseWriter, r *http.Request) {
	var user *userproto.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		respJson.WriteJSON(false, http.StatusBadRequest, "Decoding JSON failed", nil, err, w)
		//vError.WriteError("Decoding json failed!", http.StatusExpectationFailed, err, w)
	} else {
		err = validation.ValidateInputNotEmpty(user.UserName, user.UserKtp, user.UserBday, user.UserJob, user.UserEdu)

		if err != nil {
			respJson.WriteJSON(false, http.StatusBadRequest, "input error", nil, err, w)
		} else {
			err = validation.ValidateDate(user.UserBday)

			if err != nil {
				respJson.WriteJSON(false, http.StatusBadRequest, "date error", nil, err, w)
			} else {

				err = validation.ValidateKTP(user.UserKtp)

				if err != nil {
					respJson.WriteJSON(false, http.StatusBadRequest, "ktp error", nil, err, w)
				} else {
					created, err := u.service.Create(context.Background(), user)

					if err != nil {
						respJson.WriteJSON(false, http.StatusBadRequest, "Create User failed", nil, err, w)
						//vError.WriteError("Create User Failed!", http.StatusBadRequest, err, w)
					} else {
						data, err := u.service.GetByID(context.Background(), &userproto.ID{
							Id: created.UserId,
						})

						if err != nil {
							respJson.WriteJSON(false, http.StatusBadRequest, "Get User by ID failed", nil, err, w)
							//vError.WriteError("Get User by ID failed", http.StatusBadRequest, err, w)
						} else {
							respJson.WriteJSON(true, http.StatusOK, "Data created", data, nil, w)
						}
					}
				}
			}
		}
	}
}

func (u *userRoute) update(w http.ResponseWriter, r *http.Request) {
	var user *userproto.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		respJson.WriteJSON(false, http.StatusExpectationFailed, "Decoding json failed", nil, err, w)
		//vError.WriteError("Decoding json failed", http.StatusExpectationFailed, err, w)
	} else {
		err = validation.ValidateInputNotEmpty(user.UserName, user.UserKtp, user.UserBday, user.UserJob, user.UserEdu)

		if err != nil {
			respJson.WriteJSON(false, http.StatusBadRequest, "", nil, err, w)
		} else {
			err = validation.ValidateDate(user.UserBday)

			if err != nil {
				respJson.WriteJSON(false, http.StatusBadRequest, "", nil, err, w)
			} else {
				err = validation.ValidateKTP(user.UserKtp)

				if err != nil {
					respJson.WriteJSON(false, http.StatusBadRequest, "", nil, err, w)
				} else {
					id := varMux.GetVarsMux("id", r)

					data, err := u.service.GetByID(context.Background(), &userproto.ID{
						Id: id,
					})

					if err != nil {
						respJson.WriteJSON(false, http.StatusBadRequest, "Data not found", nil, err, w)
						//vError.WriteError("Get User by ID failed!", http.StatusBadRequest, err, w)
					} else {
						_, err := u.service.Update(context.Background(), &userproto.UserUpdateRequest{
							Id:   id,
							User: user,
						})

						if err != nil {
							respJson.WriteJSON(false, http.StatusBadRequest, "Updating data failed", nil, err, w)
							//vError.WriteError("Updating data failed!", http.StatusBadRequest, err, w)
						} else {
							respJson.WriteJSON(true, http.StatusOK, "Data updated", data, nil, w)
						}
					}
				}
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
		respJson.WriteJSON(false, http.StatusBadRequest, "Get User by ID failed", nil, err, w)
		//vError.WriteError("Get User By ID failed!", http.StatusBadRequest, err, w)
	} else {
		_, err := u.service.Delete(context.Background(), userID)

		if err != nil {
			respJson.WriteJSON(false, http.StatusBadRequest, "Delete User failed", nil, err, w)
			//vError.WriteError("Delete User failed!", http.StatusBadRequest, err, w)
		} else {
			respJson.WriteJSON(true, http.StatusOK, "Data deleted", data, nil, w)
		}
	}
}