package config

var (
	EtcdAddress = []string{"10.66.48.69:2379", "10.82.30.81:2379", "10.82.30.89:2379"}
	ServiceName = "com.ops.user.service.user"
)

type DBConfig struct {
	Address string `json:"address"`
	DBName string `json:dbname`
	Username string `json:username`
	Password string `json:password`
}
