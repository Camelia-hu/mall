package service

import (
	"context"
	"github.com/Camelia-hu/mall/idl/product/biz/dal/mysql"
	"github.com/Camelia-hu/mall/idl/product/module"
	product "github.com/Camelia-hu/mall/rpc_gen/kitex_gen/product"
)

type DeleteProductService struct {
	ctx context.Context
} // NewDeleteProductService new DeleteProductService
func NewDeleteProductService(ctx context.Context) *DeleteProductService {
	return &DeleteProductService{ctx: ctx}
}

// Run create note info
func (s *DeleteProductService) Run(req *product.DeleteReq) (resp *product.DeleteResp, err error) {
	// Finish your business logic.
	err = mysql.DB.Delete(&module.Product{}, req.Id).Error
	if err != nil {
		resp.Is = false
		return resp, err
	}
	resp = &product.DeleteResp{}
	resp.Is = true
	return resp, nil
}
