package builder

import (
	"log"
	"queryservice/domain/models/categories"
	"queryservice/domain/models/products"
	"queryservice/errs"

	v1 "github.com/akira-saneyoshi/store_pb/pb/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type resultBuilderImpl struct{}

func NewresultBuilderImpl() ResultBuilder {
	return &resultBuilderImpl{}
}

// *categories.Categoryを*v1.CategtoryResultに変換する
func (ins *resultBuilderImpl) BuildCategoryResult(source any) *v1.CategoryResult {
	// CategoryResultを生成する
	result := &v1.CategoryResult{Timestamp: timestamppb.Now()}
	// *categories.Categoryであるかを検証する
	if category, ok := source.(*categories.Category); ok {
		// Resultフィールドに問合せ結果を設定する
		result.Result = &v1.CategoryResult_Category{
			Category: &v1.Category{Id: category.Id(), Name: category.Name()},
		}
	} else {
		// Resultフィールドにエラーを設定する
		result.Result = &v1.CategoryResult_Error{Error: ins.BuildErrorResult(source)}
	}
	return result
}

// []*categories.Categoryを*v1.CategoriesResultに変換する
func (ins *resultBuilderImpl) BuildCategoriesResult(source any) *v1.CategoriesResult {
	// CategoriesResultを生成する
	result := &v1.CategoriesResult{Timestamp: timestamppb.Now()}
	// []categories.Category型であるかを検証する
	if categories, ok := source.([]*categories.Category); ok {
		// 問合せ結果を設定する
		c := []*v1.Category{}
		for _, category := range categories {
			c = append(c, &v1.Category{Id: category.Id(), Name: category.Name()})
		}
		result.Categories = c
	} else {
		// Errorフィールドにエラーを設定する
		result.Error = ins.BuildErrorResult(source)
	}
	return result
}

// *products.Productを*v1.ProductResultに変換する
func (ins *resultBuilderImpl) BuildProductResult(source any) *v1.ProductResult {
	// ProductResult型を生成する
	result := &v1.ProductResult{Timestamp: timestamppb.Now()}
	// *products.Productであるかを検証する
	if product, ok := source.(*products.Product); ok {
		// Resultフィールドに問合せ結果を設定する
		c := &v1.Category{Id: product.Id(), Name: product.Name()}
		result.Result = &v1.ProductResult_Product{
			Product: &v1.Product{Id: product.Id(), Name: product.Name(), Price: int32(product.Price()), Category: c},
		}
	} else {
		// Resultフィールドにエラーを設定する
		result.Result = &v1.ProductResult_Error{Error: ins.BuildErrorResult(source)}
	}
	return result
}

// []*product.Productを*v1.ProsuctsResultに変換する
func (ins *resultBuilderImpl) BuildProductsResult(source any) *v1.ProductsResult {
	// ProductsResult型を生成する
	result := &v1.ProductsResult{Timestamp: timestamppb.Now()}
	// []*products.Product型であるかを検証する
	if products, ok := source.([]*products.Product); ok {
		p := []*v1.Product{} // 問合せ結果を設定する
		for _, product := range products {
			c := &v1.Category{Id: product.Category().Id(), Name: product.Category().Name()}
			p = append(p, &v1.Product{Id: product.Id(), Name: product.Name(), Price: int32(product.Price()), Category: c})
		}
		result.Products = p
	} else {
		// Errorフィールドにエラーを設定する
		result.Error = ins.BuildErrorResult(source)
	}
	return result
}

// errs.CRUDError、errs.InternalErrorを*v1.Errorに変換する
func (ins *resultBuilderImpl) BuildErrorResult(source any) *v1.Error {
	switch v := source.(type) {
	case *errs.CRUDError:
		return &v1.Error{Type: "CRUD Error", Message: v.Error()}
	case *errs.InternalError:
		return &v1.Error{Type: "INTERNAL Error", Message: "只今、サービスを提供できません。"}
	default:
		log.Println("対応できないエラー型が指定されました。")
		return &v1.Error{Type: "INTERNAL Error", Message: "只今、サービスを提供できません。"}
	}
}
