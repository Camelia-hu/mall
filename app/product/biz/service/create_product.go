package service

import (
	"context"
	"errors"
	"github.com/Camelia-hu/mall/idl/product/biz/dal/mysql"
	"github.com/Camelia-hu/mall/idl/product/biz/dal/redis"
	"github.com/Camelia-hu/mall/idl/product/module"
	product "github.com/Camelia-hu/mall/rpc_gen/kitex_gen/product"
	"strconv"
	"strings"
)

type CreateProductService struct {
	ctx context.Context
} // NewCreateProductService new CreateProductService
func NewCreateProductService(ctx context.Context) *CreateProductService {
	return &CreateProductService{ctx: ctx}
}

// Run create note info
func (s *CreateProductService) Run(req *product.CreateReq) (resp *product.CreateResp, err error) {
	// Finish your business logic.
	if req.Name == "" || req.Categories == nil || &req.Price == nil {
		return nil, errors.New("请输入商品名称，价格以或分类标签喵～")
	}
	var Product module.Product
	Product.Name = req.Name
	Product.Price = req.Price
	Product.Categories = strings.Join(req.Categories, ",")
	Product.Picture = req.Picture
	Product.Description = req.Description
	mysql.DB.Create(&Product)
	redis.RedisClient.HSet(s.ctx, strconv.Itoa(int(Product.ID)), "name", Product.Name, "price", Product.Price, "categories", Product.Categories, "picture", Product.Picture, "description", Product.Description)
	resp = &product.CreateResp{}
	resp.Id = uint32(Product.ID)
	return resp, nil
	return
}
