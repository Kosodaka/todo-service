package transport

import (
	"context"
	"fmt"
	"github.com/Kosodaka/todo-service/internal/app"
	"github.com/Kosodaka/todo-service/internal/gRPC/proto"
	"github.com/Kosodaka/todo-service/internal/models"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

const grpcPort = 50051

type Deps struct {
}
type server struct {
	todo_service.UnimplementedTodoItemServer
}

func (s *server) GetTasks(ctx context.Context, req *todo_service.GetTasksRequest) (*todo_service.GetTasksResponse, error) {
	task := req.GetId()
	return &todo_service.GetTasksResponse{
		Item: &todo_service.ItemInfo{
			Title:       task.GetTitle(),
			Description: task.GetDescription(),
			Done:        task.GetDone(),
		},
	}, nil
}

func (s *server) UpdateTaskStatus(ctx context.Context, req *todo_service.GetTasksRequest) (*todo_service.UpdateTaskStatusResponse, error) {
	var item models.TodoItem
	status := req.GetId()
	res, err := app.Db.NewUpdate().Model(&item).Where("id=?", req.GetId()).Set("done=?", status.Done).Exec(ctx)
	if res == nil && err != nil {
		return nil, err
	}
	return &todo_service.UpdateTaskStatusResponse{
		Id:   status.GetId(),
		Done: status.GetDone()}, nil
}

func ServerGrpcConn() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Logger.Fatal().Msgf("failed to listen: %s", err.Error())
	}
	s := grpc.NewServer()
	reflection.Register(s)
	todo_service.RegisterTodoItemServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		log.Logger.Fatal().Msgf("failed to serve: %s", err.Error())
	}
}
