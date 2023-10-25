package main

import (
	"github.com/Kosodaka/todo-service/internal/app"
	transport "github.com/Kosodaka/todo-service/internal/transport/grpc"
)

func main() {
	app.Run()
	transport.ServerGrpcConn()
}
