package servers

import (
	"context"
	"log"

	pb "github.com/echo-marche/hack-tech-tips-api/proto/pb/main_api"
)

type SampleServer struct{}

func (server *SampleServer) GetSample(ctx context.Context, in *pb.SampleRequest) (*pb.SampleResponse, error) {
	log.Printf("Received: %v", in.SystemCode)
	return &pb.SampleResponse{
		Name:  "OK" + in.SystemCode,
		Email: "sample@gmail.com",
	}, nil
}
