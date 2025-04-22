package config

import "github.com/caarlos0/env/v6"

type envConfig struct {
	Addr              string `env:"RUN_ADDRESS"`
	AccrualSystemAddr string `env:"ACCRUAL_SYSTEM_ADDRESS"`
	PostgresDNS       string `env:"DATABASE_URI"`
}

func parseEnv() (*envConfig, error) {
	cfg := &envConfig{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
