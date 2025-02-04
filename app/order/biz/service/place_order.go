package service

import (
	"context"
	"github.com/Camelia-hu/mall/idl/order/biz/dal/mysql"
	"github.com/Camelia-hu/mall/idl/order/biz/dal/redis"
	"github.com/Camelia-hu/mall/idl/order/module"
	order "github.com/Camelia-hu/mall/rpc_gen/kitex_gen/order"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	// Finish your business logic.
	span := trace.SpanFromContext(s.ctx)
	var ReqOrder module.Order
	req.Address = &order.Address{}
	address := req.Address.Country + req.Address.State + req.Address.City + req.Address.StreetAddress
	ReqOrder = module.Order{
		Model:        gorm.Model{},
		Uid:          req.UserId,
		UserCurrency: req.UserCurrency,
		Address:      address,
		Email:        req.Email,
	}
	err = mysql.DB.Create(&ReqOrder).Error
	err = redis.RedisClient.HSet(s.ctx, "order:"+strconv.Itoa(int(ReqOrder.ID)), "uid", ReqOrder.Uid, "userCurrency", ReqOrder.UserCurrency, "address", ReqOrder.Address, "email", ReqOrder.Email).Err()
	err = redis.RedisClient.Expire(s.ctx, "order:"+strconv.Itoa(int(ReqOrder.ID)), 15*time.Minute).Err()
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	for _, item := range req.OrderItems {
		oneItem := module.OrderItem{
			Model:     gorm.Model{},
			OrderID:   ReqOrder.ID,
			Cost:      item.Cost,
			ProductId: item.Item.ProductId,
			Quantity:  item.Item.Quantity,
		}
		err = mysql.DB.Create(&oneItem).Error
		err = redis.RedisClient.HSet(s.ctx, "orderItem:"+strconv.Itoa(int(oneItem.OrderID)), "orderId", oneItem.OrderID, "cost", oneItem.Cost, "productId", oneItem.ProductId, "quantity", oneItem.Quantity).Err()
		err = redis.RedisClient.Expire(s.ctx, "orderItem:"+strconv.Itoa(int(oneItem.OrderID)), 15*time.Minute).Err()
		if err != nil {
			span.SetStatus(codes.Error, err.Error())
			log.Println("order item create err : ", err)
			return nil, err
		}
	}
	resp = &order.PlaceOrderResp{Order: &order.OrderResult{OrderId: strconv.Itoa(int(ReqOrder.ID))}}
	return resp, nil
}
