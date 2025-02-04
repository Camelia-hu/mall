package service

import (
	"context"
	"errors"
	"github.com/Camelia-hu/mall/idl/product/biz/dal/mysql"
	"github.com/Camelia-hu/mall/idl/product/module"
	product "github.com/Camelia-hu/mall/rpc_gen/kitex_gen/product"
	"gorm.io/gorm"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// Finish your business logic.
	var products []*module.Product
	err = mysql.DB.Where("categories like ?", "%"+req.CategoryName+"%").Find(&products).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("没有该属性喵～")
	}
	resp = &product.ListProductsResp{}
	for _, oneProduct := range products {
		respProduct := &product.Product{
			Id:          uint32(oneProduct.ID),
			Name:        oneProduct.Name,
			Description: oneProduct.Description,
			Picture:     oneProduct.Picture,
			Price:       oneProduct.Price,
		}
		resp.Products = append(resp.Products, respProduct)
	}
	return resp, nil
}
