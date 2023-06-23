package interfaces

import (
	"fmt"
	"api-gateway/pkg/config"
	"api-gateway/pkg/interfaces/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.InterfacesServiceClient
}

func InitServiceClient(c *config.Config) pb.InterfacesServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.InterfacesSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewInterfacesServiceClient(cc)
}

