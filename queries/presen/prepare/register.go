package prepare

import (
	v1 "github.com/akira-saneyoshi/store_pb/pb/v1"
	"google.golang.org/grpc"
)

// gRPCサーバの生成とQueryServiceの登録
type QueryServer struct {
	Server *grpc.Server // gRPCServer
}

// コンストラクタ
func NewQueryServer(category v1.CategoryQueryServer, product v1.ProductQueryServer) *QueryServer {
	server := grpc.NewServer()

	// CategoryQueryServerを登録する
	v1.RegisterCategoryQueryServer(server, category)
	// ProductQueryServerを登録する
	v1.RegisterProductQueryServer(server, product)
	return &QueryServer{Server: server}
}
