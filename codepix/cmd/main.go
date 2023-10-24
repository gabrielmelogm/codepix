package main

import (
	"os"

	"github.com/gabrielmelogm/codepix-go/application/grpc"
	"github.com/gabrielmelogm/codepix-go/infrastructure/db"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
