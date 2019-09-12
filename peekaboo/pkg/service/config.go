package service

import (
	"fmt"
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
type PositionConfig struct {
	X int32 `mapstructure:"x"`
	Y int32 `mapstructure:"y"`
}
type SideWindowConfig struct {
	TitleHeight int32 `mapstructure:"title_height"`
	Back PositionConfig `mapstructure:"back"`
	Home PositionConfig `mapstructure:"home"`
	Recent PositionConfig `mapstructure:"recent"`
}
type EmulatorConfig struct {
	Side []SideWindowConfig
}


var serverHostConfig ServerConfig
var sizeConfig SizeConfig
var emulatorConfig map[string]EmulatorConfig

func init() {
	if err := LoadConfig(); err != nil {
		panic(err)
	}

	emulatorConfig = make(map[string]EmulatorConfig)

	viper.GetStringMap("server")
	_ = viper.UnmarshalKey("server", &serverHostConfig)
	viper.GetStringMap("size")
	_ = viper.UnmarshalKey("size", &sizeConfig)
	var ec []SideWindowConfig
	_ = viper.UnmarshalKey("LDPlayer", &ec)
	//fmt.Fprintf(os.Stderr, "DEBUG1: %v\n", ec)
	emulatorConfig["LDPlayer"] = EmulatorConfig{
		Side: ec,
	}
	var ec2 []SideWindowConfig
	_ = viper.UnmarshalKey("Nox", &ec2)
	emulatorConfig["Nox"] = EmulatorConfig{
		Side: ec2,
	}

	//viper.Debug()
	fmt.Fprintf(os.Stderr, "%#v\n", emulatorConfig["LDPlayer"])
	fmt.Fprintf(os.Stderr, "%#v\n", emulatorConfig["Nox"])

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
func GetConfigButtonPosition(emulatorType string, titleHeight int32, command string) (position PositionConfig) {
	fmt.Fprintf(os.Stderr, "DEBUG1: %v %v %v\n", emulatorType, titleHeight, command)
	if _, ok := emulatorConfig[emulatorType]; ok {
		fmt.Fprintf(os.Stderr, "DEBUG2: %#v\n", emulatorConfig[emulatorType])
		for _, e := range emulatorConfig[emulatorType].Side {
			fmt.Fprintf(os.Stderr, "DEBUG3: %#v\n", e)
			if e.TitleHeight == titleHeight {
				fmt.Fprintf(os.Stderr, "DEBUG: %#v\n", e)
				switch command {
				case "back":
					position = e.Back
				case "home":
					position = e.Home
				case "recent":
					position = e.Recent
				}
				break
			}
		}
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