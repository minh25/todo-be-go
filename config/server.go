package config

import "fmt"

// ServerListen for specifying host & port
type ServerListen struct {
	Host string `json:"host" mapstructure:"host"`
	Port uint16 `json:"port" mapstructure:"port"`
}

func (s ServerListen) String() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

// ListenString for listen to 0.0.0.0
func (s ServerListen) ListenString() string {
	return fmt.Sprintf(":%d", s.Port)
}

// ServerConfig for configure HTTP & gRPC host & port
type ServerConfig struct {
	HTTP ServerListen `json:"http" mapstructure:"http"`
	GRPC ServerListen `json:"grpc" mapstructure:"grpc"`
}

func ServerDefaultConfig() ServerConfig {
	return ServerConfig{
		HTTP: ServerListen{
			Host: "localhost",
			Port: 10080,
		},
		GRPC: ServerListen{
			Host: "localhost",
			Port: 10443,
		},
	}
}
