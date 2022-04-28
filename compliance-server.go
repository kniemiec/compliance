package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "compliance-service/compliance"
)


var (
	port = flag.Int("port", 9999, "The server port")
)

type complianceCheckServerStruct struct {
	pb.UnimplementedComplianceCheckServer
}

func (s *complianceCheckServerStruct) CheckCompliance(ctx context.Context, in *pb.ComplianceCheckRequest) (*pb.ComplianceCheckResponse, error) {
	log.Println("request received: " + in.TransferId)
	response := &pb.ComplianceCheckResponse{}
	response.TransferId = in.TransferId
	response.ComplianceStatus = "OK"
	return response, nil
}

func (s *complianceCheckServerStruct) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Println("request received: " + in.Message)
	response := &pb.HelloResponse{}
	response.Message = "It"
	return response, nil
}

func newServer() *complianceCheckServerStruct {
	log.Println("Server created.")
	return &complianceCheckServerStruct{}
}

func main() {
	log.Println("Compliance service is starting....")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Listening on port: ", *port)
	grpcServer := grpc.NewServer()
	pb.RegisterComplianceCheckServer(grpcServer, newServer())
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
