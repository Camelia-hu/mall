package service

import (
	"context"
	"errors"
	"github.com/Camelia-hu/mall/idl/product/biz/dal/mysql"
	"github.com/Camelia-hu/mall/idl/product/module"
	product "github.com/Camelia-hu/mall/rpc_gen/kitex_gen/product"
	"gorm.io/gorm"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	// Finish your business logic.
	var Products []*module.Product
	err = mysql.DB.Where("name like ?", "%"+req.Query+"%").Find(&Products).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("没有记录喵～")
	}
	if err != nil {
		return nil, err
	}
	resp = &product.SearchProductsResp{}
	for _, oneProduct := range Products {
		respProduct := &product.Product{
			Id:          uint32(oneProduct.ID),
			Name:        oneProduct.Name,
			Description: oneProduct.Description,
			Picture:     oneProduct.Picture,
			Price:       oneProduct.Price,
		}
		resp.Results = append(resp.Results, respProduct)
	}
	return resp, nil
}
