package config

import "flag"

const (
	defaultAddr        = "localhost:8080"
	defaultPostgresDNS = "postgresql://postgres:postgres@localhost:5432/test"
)

type cmdConfig struct {
	Addr              string
	AccrualSystemAddr string
	PostgresDNS       string
}

func parseCmd() (*cmdConfig, error) {
	cfg := &cmdConfig{}

	flag.StringVar(&cfg.Addr, "a", defaultAddr, "Postgres database connection dns")
	flag.StringVar(&cfg.PostgresDNS, "d", defaultPostgresDNS, "server startup address and port")
	flag.StringVar(&cfg.AccrualSystemAddr, "r", "", "address and port of the accrual calculation system")
	flag.Parse()

	return cfg, nil
}
