package routes

import (
	"context"
	"github.com/gorilla/mux"
	userproto "github.com/vivaldy22/mekar-regis-client/proto"
	"github.com/vivaldy22/mekar-regis-client/tools/respJson"
	"net/http"
)

type jobRoute struct {
	service userproto.JobCRUDClient
}

func NewJobRoute(service userproto.JobCRUDClient, r *mux.Router) {
	handler := &jobRoute{service}

	prefJob := r.PathPrefix("/jobs").Subrouter()
	prefJob.HandleFunc("", handler.getAll).Methods(http.MethodGet)
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