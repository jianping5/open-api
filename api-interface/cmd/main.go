package main

import (
	// "fmt"
	// "log"
	// "net"

	// "api-interface/pkg/config"
	// "api-interface/pkg/pb"
	// "api-interface/pkg/services"

	"api-interface/routes"
	// "google.golang.org/grpc"
)

func main() {
	routes.R.Run("127.0.0.1:50053")
	// c, err := config.LoadConfig()

	// if err != nil {
	// 	log.Fatalln("Failed at config", err)
	// }

	// // h := db.Init(c.DBUrl)

	// lis, err := net.Listen("tcp", c.Port)

	// if err != nil {
	// 	log.Fatalln("Failed to listing:", err)
	// }

	// fmt.Println("Interfaces Svc on", c.Port)

	// s := services.Server{
	// }

	// grpcServer := grpc.NewServer()

	// pb.RegisterInterfacesServiceServer(grpcServer, &s)

	// if err := grpcServer.Serve(lis); err != nil {
	// 	log.Fatalln("Failed to serve:", err)
	// }
}
