package server

import (
	"context"
	"queryservice/domain/models/categories"
	"queryservice/presen/builder"

	v1 "github.com/akira-saneyoshi/store_pb/pb/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type categoryServer struct {
	repository categories.CategoryRepository
	builder    builder.ResultBuilder
	// v1.UnimplementedCategoryQueryServerをエンベデッドする
	v1.UnimplementedCategoryQueryServer
}

// コンストラクタ
func NewcategoryServer(repository categories.CategoryRepository, builder builder.ResultBuilder) v1.CategoryQueryServer {
	return &categoryServer{repository: repository, builder: builder}
}

// すべてのカテゴリを取得して返す
func (ins *categoryServer) List(ctx context.Context, param *emptypb.Empty) (*v1.CategoriesResult, error) {
	if categories, err := ins.repository.List(ctx); err != nil {
		return ins.builder.BuildCategoriesResult(err), nil
	} else {
		return ins.builder.BuildCategoriesResult(categories), nil
	}
}

// 指定されたIDのカテゴリを取得して返す
func (ins *categoryServer) ById(ctx context.Context, param *v1.CategoryParam) (*v1.CategoryResult, error) {
	if category, err := ins.repository.FindByCategoryId(ctx, param.GetId()); err != nil {
		return ins.builder.BuildCategoryResult(err), nil
	} else {
		return ins.builder.BuildCategoryResult(category), nil
	}
}
