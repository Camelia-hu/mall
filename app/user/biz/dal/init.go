package dal

import (
	"github.com/Camelia-hu/mall/idl/user/biz/dal/mysql"
	"github.com/Camelia-hu/mall/idl/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
