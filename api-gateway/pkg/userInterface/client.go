package userInterface

import (
	"fmt"
	"api-gateway/pkg/config"
	"api-gateway/pkg/userInterface/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.UserInterfaceServiceClient
}

func InitServiceClient(c *config.Config) pb.UserInterfaceServiceClient {

	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.UserInterfaceSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewUserInterfaceServiceClient(cc)
}

