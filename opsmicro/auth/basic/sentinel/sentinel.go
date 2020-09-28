package sentinel

import (
	"fmt"
	"github.com/alibaba/sentinel-golang/ext/datasource"
	"github.com/alibaba/sentinel-golang/ext/datasource/etcdv3"
	"github.com/micro/go-micro/v2/logger"
	"ops.was.ink/opsmicro/auth/basic/config"
	"time"
	"github.com/coreos/etcd/clientv3"
	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
)

type stateChangeTestListener struct {
}

func (s *stateChangeTestListener) OnTransformToClosed(prev circuitbreaker.State, rule circuitbreaker.Rule) {
	fmt.Printf("熔断策略: %+v, 状态: From %s to Closed, 时间: %s\n", rule.Strategy, prev.String(), time.Now().Format("2006-01-02 15:04:05"),)
}

func (s *stateChangeTestListener) OnTransformToOpen(prev circuitbreaker.State, rule circuitbreaker.Rule, snapshot interface{}) {
	fmt.Printf("熔断策略: %+v, 状态: From %s to Open, snapshot: %.2f, 时间: %s\n", rule.Strategy, prev.String(), snapshot, time.Now().Format("2006-01-02 15:04:05"),)
}

func (s *stateChangeTestListener) OnTransformToHalfOpen(prev circuitbreaker.State, rule circuitbreaker.Rule) {
	fmt.Printf("熔断策略: %+v, 状态: From %s to Half-Open, 时间: %s\n", rule.Strategy, prev.String(), time.Now().Format("2006-01-02 15:04:05"),)
}

type DatasourceGenerator struct {
	etcdv3Client *clientv3.Client
}

func NewDatasourceGenerator(config *clientv3.Config) *DatasourceGenerator {
	client, err := clientv3.New(*config)
	if err != nil {
		logger.Errorf("Fail to instance clientv3 Client, err: %+v", err)
		return nil
	}
	return &DatasourceGenerator{etcdv3Client: client}
}

func (g *DatasourceGenerator) Generate(key string, handlers ...datasource.PropertyHandler) (*etcdv3.Etcdv3DataSource, error) {
	return etcdv3.NewDatasource(g.etcdv3Client, key, handlers... )
}

func SentinelInit() {
	circuitbreaker.RegisterStateChangeListeners(&stateChangeTestListener{})
	etcdv3Gen := NewDatasourceGenerator(&clientv3.Config{
		Endpoints: config.EtcdAddress,
		DialTimeout: time.Second * 10,
	})
	if etcdv3Gen == nil {
		logger.Errorf("Fail to instance etcdv3 datasource generator.")
		return
	}
	breakerHandler := datasource.NewCircuitBreakerRulesHandler(datasource.CircuitBreakerRuleJsonArrayParser)
	breakerSource, err := etcdv3Gen.Generate("/micro/config/sentinel/rules/circuitbreaker", breakerHandler)
	if err != nil {
		fmt.Println("创建 datasource 失败...")
	}
	if err := breakerSource.Initialize(); err != nil {
		fmt.Println("初始化 datasource 失败...")
	}
}