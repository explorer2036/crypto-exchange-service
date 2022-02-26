package exchange

import (
	"context"
	"testing"

	"crypto-exchange-service/pkg/config"
)

func TestOrder(t *testing.T) {
	wrapper := NewBinanceWrapper(&config.Config{
		APIKey:    "yvoN0Vqp7blpEAzkZaL97XnINFg5C4t2X4TXZ7PnmxTkIKmnHo2EPwix3gI4NdBc",
		SecretKey: "QLSktDLG7nb9BdqKglCrHoNgHRpHaJzwPvVX7QMIrithQDz6ErjBgKvBblLBkBTa",
	})

	symbols, err := wrapper.GetSymbols(context.Background())
	if err != nil {
		t.Fatalf("query symbols: %v", err)
	}
	for _, symbol := range symbols {
		if symbol.Symbol == "LTCBTC" || symbol.Symbol == "BNBBTC" {
			id, _, err := wrapper.CreateOrder(context.Background(), symbol.Symbol, "BUY", "5", "0.005000")
			if err != nil {
				t.Fatalf("Create order: %v", err)
			}
			if err := wrapper.GetOrder(context.Background(), symbol.Symbol, id); err != nil {
				t.Fatalf("Query order %d: %v", id, err)
			}
		}
		if symbol.Symbol == "LTCUSDT" {
			id, _, err := wrapper.CreateOrder(context.Background(), symbol.Symbol, "BUY", "0.2", "125")
			if err != nil {
				t.Fatalf("Create order: %v", err)
			}
			if err := wrapper.GetOrder(context.Background(), symbol.Symbol, id); err != nil {
				t.Fatalf("Query order %d: %v", id, err)
			}
		}
	}
}
