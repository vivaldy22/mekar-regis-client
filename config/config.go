package config

import (
	"github.com/gorilla/mux"
	"github.com/vivaldy22/mekar-regis-client/middleware"
	"github.com/vivaldy22/mekar-regis-client/routes"
	"github.com/vivaldy22/mekar-regis-client/tools/viper"
	"log"
	"net/http"
)

func NewRouter() *mux.Router {
	return mux.NewRouter()
}

func RunServer(r *mux.Router) {
	port := viper.ViperGetEnv("PORT", "8080")

	log.Printf("Starting Mekar Registration API Web Server at port: %v\n", port)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}

func InitRouters(r *mux.Router) {
	r.Use(middleware.CORSMiddleware)
	r.Use(middleware.ActivityLogMiddleware)

	userClient := newUserClient()
	jobClient := newJobClient()
	eduClient := newEduClient()

	routes.NewUserRoute(userClient, r)
	routes.NewJobRoute(jobClient, r)
	routes.NewEduRoute(eduClient, r)
}