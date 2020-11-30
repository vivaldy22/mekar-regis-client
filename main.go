package main

import "github.com/vivaldy22/mekar-regis-client/config"

func main() {
	r := config.NewRouter()
	config.InitRouters(r)
	config.RunServer(r)
}