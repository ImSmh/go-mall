package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"jijizhazha1024/go-mall/common/config"
)

type Config struct {
	zrpc.RpcServerConf
	MysqlConfig config.MysqlConfig
	Consul      consul.Conf
}
