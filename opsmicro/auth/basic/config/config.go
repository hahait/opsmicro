package config

import "time"

var (
	ServiceName = "com.ops.auth.service.auth"
	TokenSecret = "gdhGcmfsrZg5kPvnhtGnFcEW6xVZKijm"
	TokenExpiredTime = time.Hour * 24
	EtcdAddress = []string{"10.66.48.69:2379", "10.82.30.81:2379", "10.82.30.89:2379"}
)