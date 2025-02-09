package builder

import v1 "github.com/akira-saneyoshi/store_pb/pb/v1"

// 実行結果をXXXResult型に変換するインターフェイス
type ResultBuilder interface {
	// *categories.Categoryを*v1.CategtoryResultに変換する
	BuildCategoryResult(source any) *v1.CategoryResult
	// []*categories.Categoryを*v1.CategoriesResultに変換する
	BuildCategoriesResult(source any) *v1.CategoriesResult
	// *products.Productを*pbProductResultに変換する
	BuildProductResult(source any) *v1.ProductResult
	// []*product.Productを*v1.ProsuctsResultに変換する
	BuildProductsResult(source any) *v1.ProductsResult
	// errs.CRUDError、errs.InternalErrorを*v1.Errorに変換する
	BuildErrorResult(source any) *v1.Error
}
