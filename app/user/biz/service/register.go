package service

import (
	"context"
	"errors"
	"github.com/Camelia-hu/mall/idl/auth/module"
	"github.com/Camelia-hu/mall/idl/user/biz/dal/mysql"
	"github.com/Camelia-hu/mall/idl/user/biz/utils"
	user "github.com/Camelia-hu/mall/rpc_gen/kitex_gen/user"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"gorm.io/gorm"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.
	span := trace.SpanFromContext(s.ctx)
	if req.Email == "" || req.Password == "" {
		span.SetStatus(codes.Error, "请输入邮箱地址或密码喵~")
		return nil, errors.New("请输入邮箱地址或密码喵~")
	}
	var usr module.User
	err1 := mysql.DB.Where("email = ?", req.Email).First(&usr).Error
	if !errors.Is(err1, gorm.ErrRecordNotFound) {
		span.SetStatus(codes.Error, "用户名已存在喵")
		return nil, errors.New("用户名已存在喵~")
	}
	if req.Password != req.ConfirmPassword {
		span.SetStatus(codes.Error, "两次密码输入不一致喵")
		return nil, errors.New("两次密码输入不一致喵~")
	}
	usr.Email = req.Email
	usr.Password = req.Password
	salt := utils.GenerateSalt()
	usr.Salt = salt
	usr.Password = utils.HashPassword(usr.Password, salt)
	mysql.DB.Create(&usr)
	var newusr module.User
	mysql.DB.Where("email = ?", usr.Email).First(&newusr)
	resp = &user.RegisterResp{}
	resp.UserId = int32(newusr.ID)
	return resp, nil
}
