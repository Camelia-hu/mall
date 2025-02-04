package dal

import (
	"github.com/Camelia-hu/mall/idl/cart/biz/dal/mysql"
	"github.com/Camelia-hu/mall/idl/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
