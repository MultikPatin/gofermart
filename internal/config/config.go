package config

import (
	"fmt"
	"go.uber.org/zap"
	"net/url"
)

type Config struct {
	Addr              string
	AccrualSystemAddr string
	PostgresDNS       *url.URL
}

func Parse(logger *zap.SugaredLogger) *Config {
	cfg := &Config{}

	envCfg, err := parseEnv()
	if err != nil {
		logger.Infow(
			"Parsed Env",
			"error", err.Error(),
		)
	}
	cmdCfg, err := parseCmd()
	if err != nil {
		logger.Infow(
			"Parsed CMD",
			"error", err.Error(),
		)
	}

	if envCfg.Addr == "" {
		cfg.Addr = cmdCfg.Addr
	} else {
		cfg.Addr = envCfg.Addr
	}
	if envCfg.AccrualSystemAddr == "" {
		cfg.AccrualSystemAddr = cmdCfg.AccrualSystemAddr
	} else {
		cfg.AccrualSystemAddr = envCfg.AccrualSystemAddr
	}
	if envCfg.PostgresDNS != "" {
		cfg.PostgresDNS, _ = parseDSN(envCfg.PostgresDNS)
	} else if cmdCfg.PostgresDNS != "" {
		cfg.PostgresDNS, _ = parseDSN(cmdCfg.PostgresDNS)
	}
	return cfg
}

func parseDSN(dsn string) (*url.URL, error) {
	parsedURL, err := url.Parse(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse DSN: %w", err)
	}

	return parsedURL, nil
}
