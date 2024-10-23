package di

import (
	"log"

	"github.com/ratheeshkumar25/opti_cut_userservice/config"
	"github.com/ratheeshkumar25/opti_cut_userservice/pkg/client/material"
	"github.com/ratheeshkumar25/opti_cut_userservice/pkg/db"
	"github.com/ratheeshkumar25/opti_cut_userservice/pkg/handlers"
	"github.com/ratheeshkumar25/opti_cut_userservice/pkg/repo"
	"github.com/ratheeshkumar25/opti_cut_userservice/pkg/server"
	"github.com/ratheeshkumar25/opti_cut_userservice/pkg/services"
)

func Init() {
	cnfg := config.LoadConfig()

	redis, err := config.SetupRedis(cnfg)
	if err != nil {
		log.Fatalf("failed to connect to redis")
	}
	twilio := config.SetupTwilio(cnfg)
	db := db.ConnectDB(cnfg)

	materialClient, err := material.ClientDial(*cnfg)
	if err != nil {
		log.Fatalf("failed to connect to material client")
	}

	userRepo := repo.NewUserRepository(db)
	userService := services.NewUserService(userRepo, redis, twilio, materialClient)
	userHandler := handlers.NewUserHandler(userService)

	err = server.NewGrpcUserServer(cnfg.GrpcPort, userHandler)
	if err != nil {
		log.Fatalf("failed to start gRPC server %v", err.Error())
	}

}
