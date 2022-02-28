package config

import (
	"os"
	"strconv"

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
}

// New returns the server config
func New() *Config {
	settings := Config{
		HTTPAddr:     ":9000",
		GRPCAddr:     ":9001",
		ReadTimeout:  5,
		WriteTimeout: 10,
		APIKey:       "az6j45YMUmZbkTGJoez2pGUvHEAJeF21BwcAoQBUSecF5RYBTiyeqDjDPbZmE04y",
		SecretKey:    "ZANmzkuuVCpemKJGnQ2Oi15NSeVDKPEG7tp8MvVlPY2kGJBZbHhxp8fGUdOJkl1u",
	}

	if addr := os.Getenv("HTTP_ADDR"); addr != "" {
		settings.HTTPAddr = addr
	}
	if addr := os.Getenv("GRPC_ADDR"); addr != "" {
		settings.GRPCAddr = addr
	}
	if read := os.Getenv("READ_TIMEOUT"); read != "" {
		timeout, err := strconv.Atoi(read)
		if err == nil {
			settings.ReadTimeout = timeout
		}
	}
	if write := os.Getenv("WRITE_TIMEOUT"); write != "" {
		timeout, err := strconv.Atoi(write)
		if err == nil {
			settings.WriteTimeout = timeout
		}
	}

	logrus.Infof("server config: %v", settings)
	return &settings
}
