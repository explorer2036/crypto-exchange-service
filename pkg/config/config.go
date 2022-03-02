package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Config structure for server
type Config struct {
	HTTPAddr     string `json:"http_address"`
	GRPCAddr     string `json:"grpc_address"`
	ReadTimeout  int    `json:"read_timeout"`
	WriteTimeout int    `json:"write_timeout"`
	APIKey       string `json:"api_key"`
	SecretKey    string `json:"secret_key"`
	TestNetwork  string `json:"test_network"`
}

// New returns the server config
func New() *Config {
	settings := Config{
		HTTPAddr:     ":9000",
		GRPCAddr:     ":9001",
		ReadTimeout:  5,
		WriteTimeout: 10,
		APIKey:       "pYD6avUt6U4Q8KRzH6NayZ3GjjtakWUPXVyoRnqiiAmyrvxcr1wUbWs0ZCbok2Hg",
		SecretKey:    "0VMAEaLSbpevMaaJVwA58kuftzvrpx71i7XEncVnBy4KaqshInfWHNhvO6ch4K9d",
		TestNetwork:  "yes",
	}

	if addr := os.Getenv("HTTP_ADDR"); addr != "" {
		settings.HTTPAddr = addr
	}
	if addr := os.Getenv("GRPC_ADDR"); addr != "" {
		settings.GRPCAddr = addr
	}
	if key := os.Getenv("API_KEY"); key != "" {
		settings.APIKey = key
	}
	if key := os.Getenv("SECRET_KEY"); key != "" {
		settings.SecretKey = key
	}
	if network := os.Getenv("TEST_NETWORK"); network != "" {
		settings.TestNetwork = network
	}

	logrus.Infof("server config: %v", settings)
	return &settings
}
