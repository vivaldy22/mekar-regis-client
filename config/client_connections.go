package config

import (
	userproto "github.com/vivaldy22/mekar-regis-client/proto"
	"github.com/vivaldy22/mekar-regis-client/tools/viper"
	"google.golang.org/grpc"
	"log"
)

func newUserClient() userproto.UserCRUDClient {
	host := viper.ViperGetEnv("GRPC_USER_HOST", "localhost")
	port := viper.ViperGetEnv("GRPC_USER_PORT", "1010")
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return userproto.NewUserCRUDClient(conn)
}

func newJobClient() userproto.JobCRUDClient {
	host := viper.ViperGetEnv("GRPC_USER_HOST", "localhost")
	port := viper.ViperGetEnv("GRPC_USER_PORT", "1010")
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return userproto.NewJobCRUDClient(conn)
}

func newEduClient() userproto.EduCRUDClient {
	host := viper.ViperGetEnv("GRPC_USER_HOST", "localhost")
	port := viper.ViperGetEnv("GRPC_USER_PORT", "1010")
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return userproto.NewEduCRUDClient(conn)
}