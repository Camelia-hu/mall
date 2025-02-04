package service

import (
	"context"
	"errors"
	"github.com/Camelia-hu/mall/idl/product/biz/dal/mysql"
	"github.com/Camelia-hu/mall/idl/product/biz/dal/redis"
	"github.com/Camelia-hu/mall/idl/product/module"
	product "github.com/Camelia-hu/mall/rpc_gen/kitex_gen/product"
	Redis "github.com/redis/go-redis/v9"
	"strconv"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// Finish your business logic.
	var Product module.Product
	result, err := redis.RedisClient.HGetAll(s.ctx, strconv.Itoa(int(req.Id))).Result()
	if err == nil {
		price, _ := strconv.Atoi(result["price"])
		pro := product.Product{
			Id:          req.Id,
			Name:        result["name"],
			Description: result["description"],
			Picture:     result["picture"],
			Price:       float32(price),
		}
		resp = &product.GetProductResp{Product: &pro}
		return resp, nil
	}
	if errors.Is(err, Redis.Nil) {
		err = mysql.DB.Where("id = ?", req.Id).First(&Product).Error
		if err != nil {
			return nil, err
		}
		redis.RedisClient.HSet(s.ctx, strconv.Itoa(int(Product.ID)), "name", Product.Name, "price", Product.Price, "categories", Product.Categories, "picture", Product.Picture, "description", Product.Description)
		pro := product.Product{
			Id:          uint32(Product.ID),
			Name:        Product.Name,
			Description: Product.Description,
			Picture:     Product.Picture,
			Price:       Product.Price,
		}
		resp = &product.GetProductResp{Product: &pro}
		return resp, nil
	}
	return nil, err
}
