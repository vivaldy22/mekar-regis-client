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

type jobRoute struct {
	service userproto.JobCRUDClient
}

func NewJobRoute(service userproto.JobCRUDClient, r *mux.Router) {
	handler := &jobRoute{service}

	prefJob := r.PathPrefix("/jobs").Subrouter()
	prefJob.Use(middleware.AdminJWTMiddleware.Handler)

	prefJob.HandleFunc("", handler.getAll).Methods(http.MethodGet)
	prefJob.HandleFunc("/{id}", handler.getByID).Methods(http.MethodGet)
}

func (j *jobRoute) getAll(w http.ResponseWriter, r *http.Request) {
	data, err := j.service.GetAll(context.Background(), new(userproto.Empty))

	if err != nil {
		respJson.WriteJSON(false, http.StatusBadRequest, "Get All Jobs failed", nil, err, w)
		//vError.WriteError("Get All Jobs failed!", http.StatusBadRequest, err, w)
	} else {
		respJson.WriteJSON(true, http.StatusOK, "Data found", data.List, nil, w)
	}
}

func (j *jobRoute) getByID(w http.ResponseWriter, r *http.Request) {
	id := varMux.GetVarsMux("id", r)

	data, err := j.service.GetByID(context.Background(), &userproto.ID{
		Id: id,
	})

	if err != nil {
		respJson.WriteJSON(false, http.StatusBadRequest, "Get Job by ID failed", nil, err, w)
		//vError.WriteError("Get User by ID failed", http.StatusBadRequest, err, w)
	} else {
		respJson.WriteJSON(true, http.StatusOK, "Data found", data, nil, w)
	}
}