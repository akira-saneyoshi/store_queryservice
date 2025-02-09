package interceptor

import (
	"context"
	"regexp"

	v1 "github.com/akira-saneyoshi/store_pb/pb/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UUID形式かをチェックする
func isUUID(u string) bool {
	var uuidPattern = regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`)
	return uuidPattern.MatchString(u)
}

// UUID形式をチェックするインターセプタ
func UUIDValidationInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	var id string
	switch info.FullMethod {
	case "/proto.CategoryQuery/ById":
		param, _ := req.(*v1.CategoryParam)
		id = param.Id
	case "/proto.ProductQuery/ById":
		param, _ := req.(*v1.ProductParam)
		id = param.Id
	}
	if id != "" {
		if !isUUID(id) { // UUIDかを検証
			return nil, status.Error(codes.InvalidArgument, "UUID形式ではありません。")
		}
	}
	return handler(ctx, req)
}
