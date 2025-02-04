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

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	// Finish your business logic.
	span := trace.SpanFromContext(s.ctx)
	var carts []module.Cart
	result, err := redis.RedisClient.HGetAll(s.ctx, strconv.Itoa(int(req.UserId))).Result()
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		err = mysql.DB.Where("uid = ?", req.UserId).Find(&carts).Error
		if err != nil {
			span.SetStatus(codes.Error, err.Error())
			log.Println("get cart err : ", err)
			return nil, err
		}
		resp = &cart.GetCartResp{
			Cart: &cart.Cart{
				UserId: req.UserId,
				Items:  []*cart.CartItem{},
			},
		}
		for _, oneCart := range carts {
			item := &cart.CartItem{
				ProductId: uint32(oneCart.ProductId),
				Quantity:  int32(oneCart.Quantity),
			}
			resp.Cart.Items = append(resp.Cart.Items, item)
		}
		return resp, nil
	}
	resp = &cart.GetCartResp{
		Cart: &cart.Cart{
			UserId: req.UserId,
			Items:  []*cart.CartItem{},
		},
	}
	for key, value := range result {
		var oneItem *cart.CartItem
		ikey, _ := strconv.Atoi(key)
		oneItem.ProductId = uint32(ikey)
		ivalue, _ := strconv.Atoi(value)
		oneItem.Quantity = int32(ivalue)
		resp.Cart.Items = append(resp.Cart.Items, oneItem)
	}
	return resp, nil
}
