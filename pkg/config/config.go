package config

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

// Config structure for server
type Config struct {
	Address      string `json:"address"`
	ReadTimeout  int    `json:"read_timeout"`
	WriteTimeout int    `json:"write_timeout"`
	APIKey       string `json:"api_key"`
	SecretKey    string `json:"secret_key"`
}

// New returns the server config
func New() *Config {
	settings := Config{
		Address:      ":9000",
		ReadTimeout:  5,
		WriteTimeout: 10,
		APIKey:       "yvoN0Vqp7blpEAzkZaL97XnINFg5C4t2X4TXZ7PnmxTkIKmnHo2EPwix3gI4NdBc",
		SecretKey:    "QLSktDLG7nb9BdqKglCrHoNgHRpHaJzwPvVX7QMIrithQDz6ErjBgKvBblLBkBTa",
	}

	if address := os.Getenv("SERVER_ADDR"); address != "" {
		settings.Address = address
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
