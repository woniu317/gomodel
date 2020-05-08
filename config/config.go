package config

import (
	"sync"
	"time"
)

type Mysql struct {
	Server          string
	MaxIdle         int
	MaxOpen         int
	ConnMaxLifetime time.Duration
}

var (
	MysqlStuConfig     Mysql
	MysqlTeacherConfig Mysql
	mysqlOnce          sync.Once
)

func init() {
	mysqlOnce.Do(func() {
		// 完成mysql配置初始化
	})
}
