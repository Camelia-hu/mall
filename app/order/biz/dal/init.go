package dal

import (
	"github.com/Camelia-hu/mall/idl/order/biz/dal/mysql"
	"github.com/Camelia-hu/mall/idl/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
