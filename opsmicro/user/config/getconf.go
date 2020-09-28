package config

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/etcd"
	"time"
)

var (
	conf config.Config
	err error
)

func ConfInit() error {
	etcdsource := etcd.NewSource(
		etcd.WithAddress(EtcdAddress...),
		// 完整的节点名称: /micro/config/<微服务名>/<配置名>/<环境名>/<key>: '<json格式的数据>'
		// eg. /micro/config/user/user_db/online: '{}'
		etcd.WithPrefix("/micro/config/user"),
		etcd.StripPrefix(true),
		etcd.WithDialTimeout(time.Second * 2),
	)

	if conf, err = config.NewConfig(config.WithSource(etcdsource)); err != nil {
		return err
	}
	return nil
}

func GetConf() config.Config {
	return conf
}