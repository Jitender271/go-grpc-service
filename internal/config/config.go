package config

import "time"

type AppConfig struct {
	LogLevel              string        `mapstructure:"LOG_LEVEL"`
	GRPCPort              string        `mapstructure:"GRPC_PORT"`
	ReverseProxyHttpPort  string        `mapstructure:"REVERSE_PROXY_PORT"`
	HttpReadTimeout       time.Duration `mapstructure:"HTTP_READ_TIMEOUT"`
	HttpWriteTimeout      time.Duration `mapstructure:"HTTP_WRITE_TIMEOUT"`
	HttpIdleTimeout       time.Duration `mapstructure:"HTTP_IDLE_TIMEOUT"`
	GRPCConnectionTimeout time.Duration `mapstructure:"GRPC_CONNECTION_TIMEOUT"`
	DbConfigs             DbConfigs     `mapstructure:"DB_CONFIGS"`
}

type DbConfigs struct {
	DBHosts                  []string      `mapstructure:"DB_HOSTS"`
	DBUsername               string        `mapstructure:"DB_USERNAME"`
	DBPassword               string        `mapstructure:"DB_PASSWORD"`
	DBPort                   int           `mapstructure:"DB_PORT"`
	DBKeyspace               string        `mapstructure:"DB_KEYSPACE"`
	DBConsistency            string        `mapstructure:"DB_CONSISTENCY"`
	DBConnectionsPerHost     int           `mapstructure:"DB_CONNECTIONS_PER_HOST"`
	DBConnectTimeout         time.Duration `mapstructure:"DB_CONNECT_TIMEOUT"`
	DBWriteTimeout           time.Duration `mapstructure:"DB_WRITE_TIMEOUT"`
	DBReadTimeout            time.Duration `mapstructure:"DB_READ_TIMEOUT"`
	DBKeepAliveTime          time.Duration `mapstructure:"DB_KEEP_ALIVE_TIME"`
	DBReConnectionMaxRetries int           `mapstructure:"DB_RECONNECTION_MAX_RETRIES"`
	DBReConnectionInterval   time.Duration `mapstructure:"DB_RECONNECTION_INTERVAL"`
}
