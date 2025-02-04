package dal

import (
	"github.com/Camelia-hu/mall/idl/auth/biz/dal/mysql"
	"github.com/Camelia-hu/mall/idl/auth/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
