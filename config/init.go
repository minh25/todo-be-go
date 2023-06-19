package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

// Config for app configuration
type Config struct {
	Server ServerConfig `json:"server" mapstructure:"server"`
	MySQL  MySQLConfig  `json:"mysql" mapstructure:"mysql"`
}

var Instance *Config

func init() {
	var err error
	Instance, err = Load()
	if err != nil {
		panic(err)
	}
	fmt.Println(Instance)
}

// Load system env config
func Load() (*Config, error) {
	// You should set default config value here
	c := &Config{
		MySQL:  MySQLDefaultConfig(),
		Server: ServerDefaultConfig(),
	}

	// --- hacking to load reflect structure config into env ----//
	viper.SetConfigType("json")
	configBuffer, err := json.Marshal(c)

	if err != nil {
		return nil, err
	}

	err = viper.ReadConfig(bytes.NewBuffer(configBuffer))
	if err != nil {
		return nil, err
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))

	// -- end of hacking --//
	viper.AutomaticEnv()
	err = viper.Unmarshal(c)
	return c, err
}
