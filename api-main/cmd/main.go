package main

import (
	"api-main/config"
	"api-main/inner"
	"api-main/pb"
	"api-main/routes"
	// "fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// http 服务
	go func() {
		err := routes.R.Run(config.Config.Gin.Address)
		if err != nil {
			return
		}
	}()


	// grpc 服务
	lis, _ := net.Listen("tcp", ":50051")

	s := inner.Server{}
	grpcServer := grpc.NewServer()
	pb.RegisterUserInterfaceServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
