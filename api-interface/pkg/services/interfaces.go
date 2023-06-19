package services

import (
	"context"
	"net/http"

	"api-interface/pkg/pb"
)

type Server struct {
	pb.UnimplementedInterfacesServiceServer
}

func (s *Server) GetName(ctx context.Context, req *pb.GetNameRequest) (*pb.GetNameResponse, error) {


	return &pb.GetNameResponse{
		Status: http.StatusCreated,
		Data: "Tom",
	}, nil
}