package config

import (
	"github.com/bells307/adv-service/pkg/mongodb"
	"github.com/spf13/viper"
)

type Config struct {
	HttpListen string                `mapstructure:"HTTP_LISTEN"`
	GrpcListen string                `mapstructure:"GRPC_LISTEN"`
	MongoDB    mongodb.MongoDBConfig `mapstructure:"MONGODB"`
}

func LoadConfig(path string) (cfg Config, err error) {
	viper.SetConfigFile(path)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&cfg)
	return
}
