package material

import (
	"fmt"
	"log"

	"github.com/ratheeshkumar25/opti_cut_userservice/config"
	pb "github.com/ratheeshkumar25/opti_cut_userservice/pkg/client/material/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ClientDial method connects to the service to the materialservice client
func ClientDial(cfg config.Config) (pb.MaterialServiceClient, error) {
	grpcAddr := fmt.Sprintf("material-service:%s", cfg.MateialPort)
	grpc, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error dialing to grpc to client : %s", err.Error())
		return nil, err
	}
	log.Printf("Successfully connected to material client at port : %s", cfg.MateialPort)
	return pb.NewMaterialServiceClient(grpc), nil
}
