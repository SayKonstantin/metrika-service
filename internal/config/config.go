package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type GRPC struct {
	Ip   string `yaml:"grpc_ip" env:"GRPC_IP"`
	Port int    `yaml:"grpc_port" env:"GRPC_PORT"`
}

type TG struct {
	IsEnabled bool   `yaml:"is_enabled" env:"TG_ENABLED"`
	Token     string `yaml:"token" env:"TG_TOKEN"`
	Chat      int64  `yaml:"chat" env:"TG_CHAT"`
}

type ServerConfig struct {
	KeysDir        string `yaml:"keys_dir" env:"KEYS_DIR"`
	GRPC           `yaml:"grpc"`
	TG             `yaml:"tg"`
	PrometheusAddr string `yaml:"prometheus_addr" env:"PROMETHEUS_ADDR"`
	AttachmentsDir string `yaml:"attachments_dir" env:"ATTACHMENTS_DIR"`
	Yandex         `yaml:"yandex"`
}

type Yandex struct {
	Token string `yaml:"yandex_token" env:"YANDEX_TOKEN"`
}

func NewServerConfig(filePath string, useEnv bool) (*ServerConfig, error) {
	cfg := &ServerConfig{}

	if useEnv {
		err := cleanenv.ReadEnv(cfg)
		if err != nil {
			return nil, fmt.Errorf("read ENV error: %w", err)
		}
	} else {
		err := cleanenv.ReadConfig(filePath, cfg)
		if err != nil {
			return nil, fmt.Errorf(" read config file error: %w", err)
		}
	}

	return cfg, nil
}
