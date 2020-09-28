package config

import (
	"fmt"
)

var (
	DBConf = new(DBConfig)
)

func init() {
	// 1. 初始化 Etcd Config 对象, 用于从 etcd 中获取配置信息
	if err := ConfInit(); err != nil {
		panic(fmt.Sprintf("etcd config 初始化失败, 错误信息: %s", err.Error()))
	}

	// 2. 获取 gorm 配置信息; key: /micro/config/user/online/userdb
	if err := conf.Get("online", "userdb").Scan(DBConf); err != nil {
	//if err := conf.Get("micro", "config", "user", "online", "userdb").Scan(DBConf); err != nil {
		panic(fmt.Sprintf("从 etcd 中获取 gorm 配置信息出错, 错误信息: %s", err.Error()))
	}

	watcher, err := conf.Watch()
	if err != nil {
		panic(fmt.Sprintf("向 etcd 注册一个 /micro/config/user/online/userdb 的 watcher 失败, 错误信息: %s", err.Error()))
	}
	go func() {
		fmt.Println("现在开始注册一个 watcher...")
		for {
			fmt.Println("我在 goroutine 里面执行...")
			v, err := watcher.Next()
			if err != nil {
				panic(fmt.Sprintf("从 etcd watcher 中获取 Value 对象失败, 错误信息: %s", err.Error()))
			}

			fmt.Println("当前 config.Watch 从 etcd 中获取的 数据: ", string(v.Bytes()))
		}
	}()
}
