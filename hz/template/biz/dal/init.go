package dal

import (
	"github.com/hertz/hello/biz/dal/mysql"
	"github.com/hertz/hello/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
