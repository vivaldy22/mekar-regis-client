package routes

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/vivaldy22/mekar-regis-client/middleware"
	userproto "github.com/vivaldy22/mekar-regis-client/proto"
	"github.com/vivaldy22/mekar-regis-client/tools/respJson"
	"github.com/vivaldy22/mekar-regis-client/tools/varMux"
	"net/http"
)

type eduRoute struct {
	service userproto.EduCRUDClient
}

func NewEduRoute(service userproto.EduCRUDClient, r *mux.Router) {
	handler := &eduRoute{service}

	prefEdu := r.PathPrefix("/edus").Subrouter()
	prefEdu.Use(middleware.AdminJWTMiddleware.Handler)

	prefEdu.HandleFunc("", handler.getAll).Methods(http.MethodGet)
	prefEdu.HandleFunc("/{id}", handler.getByID).Methods(http.MethodGet)
}

func (e *eduRoute) getAll(w http.ResponseWriter, r *http.Request) {
	data, err := e.service.GetAll(context.Background(), new(userproto.Empty))

	if err != nil {
		respJson.WriteJSON(false, http.StatusBadRequest, "Get All Edu failed", nil, err, w)
		//vError.WriteError("Get All Edu failed!", http.StatusBadRequest, err, w)
	} else {
		respJson.WriteJSON(true, http.StatusOK, "Data found", data.List, nil, w)
	}
}

func (e *eduRoute) getByID(w http.ResponseWriter, r *http.Request) {
	id := varMux.GetVarsMux("id", r)

	data, err := e.service.GetByID(context.Background(), &userproto.ID{
		Id: id,
	})

	if err != nil {
		respJson.WriteJSON(false, http.StatusBadRequest, "Get Edu by ID failed", nil, err, w)
		//vError.WriteError("Get User by ID failed", http.StatusBadRequest, err, w)
	} else {
		respJson.WriteJSON(true, http.StatusOK, "Data found", data, nil, w)
	}
}