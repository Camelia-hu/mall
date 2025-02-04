package dal

import (
	"github.com/Camelia-hu/mall/idl/payment/biz/dal/mysql"
	"github.com/Camelia-hu/mall/idl/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
