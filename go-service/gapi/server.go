package gapi

import (
	"fmt"

	db "github.com/gbengafagbola/microservice/go-service/db/sqlc"
	"github.com/gbengafagbola/microservice/go-service/pb"
	"github.com/gbengafagbola/microservice/go-service/token"
	"github.com/gbengafagbola/microservice/go-service/util"
	"github.com/gbengafagbola/microservice/go-service/worker"
)

// Server serves gRPC requests for our delivery service.
type Server struct {
	pb.UnimplementedGoServiceServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	taskDistributor worker.TaskDistributor
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
		taskDistributor: taskDistributor, 
	}

	return server, nil
}
