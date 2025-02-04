package service

import (
	"context"
	"errors"
	"github.com/Camelia-hu/mall/idl/user/biz/dal/mysql"
	"github.com/Camelia-hu/mall/idl/user/biz/utils"
	"github.com/Camelia-hu/mall/idl/user/module"
	user "github.com/Camelia-hu/mall/rpc_gen/kitex_gen/user"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.
	span := trace.SpanFromContext(s.ctx)
	var usr module.User
	if req.Password == "" || req.Email == "" {
		span.SetStatus(codes.Error, "请输入邮箱地址或密码喵")
		return nil, errors.New("请输入邮箱地址或密码喵～")
	}
	err = mysql.DB.Where("email = ?", req.Email).First(&usr).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		span.SetStatus(codes.Error, "该用户名不存在喵")
		return nil, errors.New("该用户名不存在喵～")
	}
	if utils.HashPassword(req.Password, usr.Salt) != usr.Password {
		span.SetStatus(codes.Error, "密码输入错误喵～")
		return nil, errors.New("密码输入错误喵～")
	}
	resp = &user.LoginResp{}
	resp.UserId = int32(usr.ID)
	return resp, nil
	return
}
