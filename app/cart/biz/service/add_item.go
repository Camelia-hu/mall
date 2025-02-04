package service

import (
	"context"
	"github.com/Camelia-hu/mall/idl/cart/biz/dal/mysql"
	"github.com/Camelia-hu/mall/idl/cart/biz/dal/redis"
	"github.com/Camelia-hu/mall/idl/cart/module"
	cart "github.com/Camelia-hu/mall/rpc_gen/kitex_gen/cart"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"log"
	"strconv"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	// Finish your business logic.
	span := trace.SpanFromContext(s.ctx)
	var addCart module.Cart
	addCart.Uid = uint(req.UserId)
	addCart.ProductId = uint(req.Item.ProductId)
	addCart.Quantity = int(req.Item.Quantity)
	exist, _ := redis.RedisClient.HExists(s.ctx, strconv.Itoa(int(req.UserId)), strconv.Itoa(int(req.Item.ProductId))).Result()
	if exist {
		redis.RedisClient.HIncrBy(s.ctx, strconv.Itoa(int(req.UserId)), strconv.Itoa(int(req.Item.ProductId)), int64(req.Item.Quantity))
	} else {
		redis.RedisClient.HSet(s.ctx, strconv.Itoa(int(req.UserId)), strconv.Itoa(int(req.Item.ProductId)), req.Item.Quantity)
	}
	err = mysql.DB.Create(&addCart).Error
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		log.Println("add item err : ", err)
		return nil, err
	}
	resp = &cart.AddItemResp{}
	return resp, nil
}
