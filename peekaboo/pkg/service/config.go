package service

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

type HostsConfig struct {
	HttpHosts string `mapstructure:"http_hosts"`
	GrpcHosts string `mapstructure:"grpc_hosts"`
}
type ServerConfig struct {
	Pikabu HostsConfig `mapstructure:"pikabu"`
}
type SizeConfig struct {
	Width  int32 `mapstructure:"width"`
	Height int32 `mapstructure:"height"`
}

var serverHostConfig ServerConfig
var sizeConfig SizeConfig

func init() {
	if err := LoadConfig(); err != nil {
		panic(err)
	}

	viper.GetStringMap("server")
	_ = viper.UnmarshalKey("server", &serverHostConfig)
	viper.GetStringMap("size")
	_ = viper.UnmarshalKey("size", &sizeConfig)

	//viper.Debug()
}
func LoadConfig() (err error) {
	viper.SetConfigFile(os.Getenv("PIKABU_HOME") + "/config/config.json")
	if err = viper.ReadInConfig(); err != nil {
		viper.SetConfigFile(os.Getenv("PIKABU_HOME") + "/worker-peekaboo" + "/config/config.json")
		if err = viper.ReadInConfig(); err != nil {
			return
		}
		return
	}
	return
}
func GetConfigServerHttp() string {
	if strings.Contains(serverHostConfig.Pikabu.HttpHosts, "PORT") {
		port := "17090"
		hosts := strings.Replace(serverHostConfig.Pikabu.HttpHosts, "PORT", port, 1)
		return hosts
	} else {
		return serverHostConfig.Pikabu.HttpHosts
	}
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
func GetConfigSizeWidth() int32 {
	return sizeConfig.Width
}
func GetConfigSizeHeight() int32 {
	return sizeConfig.Height
}