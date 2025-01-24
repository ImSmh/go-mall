package svc

import (
	"jijizhazha1024/go-mall/services/users/internal/config"
	"jijizhazha1024/go-mall/services/users/internal/db"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Mysql  sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := db.NewMysql(c.MysqlConfig)
	return &ServiceContext{
		Config: c,
		Mysql:  mysql,
	}
}
