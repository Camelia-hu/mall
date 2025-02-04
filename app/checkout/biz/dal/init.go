package dal

import (
	"github.com/Camelia-hu/mall/idl/checkout/biz/dal/mysql"
	"github.com/Camelia-hu/mall/idl/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
