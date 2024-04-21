package server

import (
	"context"

	"github.com/felipemalacarne/lumina/logger/api/models"
	pb "github.com/felipemalacarne/lumina/logger/proto"
	"go.mongodb.org/mongo-driver/bson"
)

type LoggerServer struct {
	pb.UnimplementedLoggerServer
	logger *models.Logger
}

func NewLoggerServer(logger *models.Logger) *LoggerServer {
	return &LoggerServer{logger: logger}
}

// Log implements the Log RPC method.
func (s *LoggerServer) Log(ctx context.Context, req *pb.LogRequest) (*pb.LogResponse, error) {
	data := bson.M{}
	for k, v := range req.Data {
		data[k] = v
	}

	entry := s.logger.NewEntry(
		models.Level(req.Level),
		req.Message,
		data,
		models.Service(req.Service),
	)

	err := s.logger.Log(entry)
	if err != nil {
		return &pb.LogResponse{Success: false}, err
	}

	return &pb.LogResponse{Success: true}, nil
}
