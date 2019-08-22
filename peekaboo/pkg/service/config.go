package service

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

type HostsConfig struct {
	GrpcHosts string `mapstructure:"grpc_hosts"`
}
type ServerConfig struct {
	Pikabu HostsConfig `mapstructure:"pikabu"`
}

var serverHostConfig ServerConfig

func init() {
	if err := LoadConfig(); err != nil {
		panic(err)
	}

	viper.GetStringMap("server")
	_ = viper.UnmarshalKey("server", &serverHostConfig)

	//viper.Debug()
}
func LoadConfig() (err error) {
	viper.SetConfigFile(os.Getenv("PIKABU_HOME") + "/config/config.json")
	if err = viper.ReadInConfig(); err != nil {
		viper.SetConfigFile(os.Getenv("PIKABU_HOME") + "/peekaboo" + "/config/config.json")
		if err = viper.ReadInConfig(); err != nil {
			return
		}
		return
	}
	return
}
func GetConfigServerGrpc() string {
	if strings.Contains(serverHostConfig.Pikabu.GrpcHosts, "PORT") {
		port := "17091"
		hosts := strings.Replace(serverHostConfig.Pikabu.GrpcHosts, "PORT", port, 1)
		return hosts
	} else {
		return serverHostConfig.Pikabu.GrpcHosts
	}
}