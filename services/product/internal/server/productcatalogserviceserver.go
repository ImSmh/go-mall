// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.5
// Source: product.proto

package server

import (
	"context"

	"jijizhazha1024/go-mall/services/product/internal/logic"
	"jijizhazha1024/go-mall/services/product/internal/svc"
	"jijizhazha1024/go-mall/services/product/product"
)

type ProductCatalogServiceServer struct {
	svcCtx *svc.ServiceContext
	product.UnimplementedProductCatalogServiceServer
}

func NewProductCatalogServiceServer(svcCtx *svc.ServiceContext) *ProductCatalogServiceServer {
	return &ProductCatalogServiceServer{
		svcCtx: svcCtx,
	}
}

// 根据商品id得到商品详细信息
func (s *ProductCatalogServiceServer) GetProduct(ctx context.Context, in *product.GetProductReq) (*product.GetProductResp, error) {
	l := logic.NewGetProductLogic(ctx, s.svcCtx)
	return l.GetProduct(in)
}

// 添加新商品
func (s *ProductCatalogServiceServer) CreateProduct(ctx context.Context, in *product.CreateProductReq) (*product.CreateProductResp, error) {
	l := logic.NewCreateProductLogic(ctx, s.svcCtx)
	return l.CreateProduct(in)
}

// 修改商品
func (s *ProductCatalogServiceServer) UpdateProduct(ctx context.Context, in *product.UpdateProductReq) (*product.UpdateProductResp, error) {
	l := logic.NewUpdateProductLogic(ctx, s.svcCtx)
	return l.UpdateProduct(in)
}

// 删除商品
func (s *ProductCatalogServiceServer) DeleteProduct(ctx context.Context, in *product.DeleteProductReq) (*product.DeleteProductResp, error) {
	l := logic.NewDeleteProductLogic(ctx, s.svcCtx)
	return l.DeleteProduct(in)
}

// 分页得到全部商品
func (s *ProductCatalogServiceServer) GetAllProduct(ctx context.Context, in *product.GetAllProductsReq) (*product.GetAllProductsResp, error) {
	l := logic.NewGetAllProductLogic(ctx, s.svcCtx)
	return l.GetAllProduct(in)
}
