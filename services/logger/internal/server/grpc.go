package server

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	pb "github.com/felipemalacarne/lumina/logger/cmd/api/v1"
	"github.com/felipemalacarne/lumina/logger/internal/database"
	"github.com/felipemalacarne/lumina/logger/internal/logger"
)

type LoggerServer struct {
	pb.UnimplementedLoggerServer
	logger *logger.Logger
}

func NewLoggerServer(logger *logger.Logger) *LoggerServer {
	return &LoggerServer{logger: logger}
}

// Log implements the Log RPC method.
func (s *LoggerServer) Log(ctx context.Context, req *pb.LogRequest) (*pb.LogResponse, error) {
	data := bson.M{}
	for k, v := range req.Data {
		data[k] = v
	}

	entry := s.logger.NewEntry(
		database.Level(req.Level),
		req.Message,
		data,
		database.Service(req.Service),
	)

	err := s.logger.Log(entry)
	if err != nil {
		return &pb.LogResponse{Success: false}, err
	}

	return &pb.LogResponse{Success: true}, nil
}
