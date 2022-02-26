package exchange

import (
	"context"
	"encoding/json"

	"crypto-exchange-service/pkg/config"

	"github.com/adshao/go-binance/v2"
	"github.com/sirupsen/logrus"
)

// BinanceWrapper represents the wrapper for the Binance exchange.
type BinanceWrapper struct {
	client *binance.Client
}

// NewBinanceWrapper creates a generic wrapper of the binance API.
func NewBinanceWrapper(settings *config.Config) *BinanceWrapper {
	binance.UseTestnet = true
	client := binance.NewClient(settings.APIKey, settings.SecretKey)
	return &BinanceWrapper{
		client: client,
	}
}

func (s *BinanceWrapper) CreateOrder(ctx context.Context, symbol string, side string, quantity string, price string) (int64, int64, error) {
	res, err := s.client.NewCreateOrderService().
		Symbol(symbol).
		Side(binance.SideType(side)).
		Type(binance.OrderTypeLimit).
		TimeInForce(binance.TimeInForceTypeGTC).
		Quantity(quantity).
		Price(price).
		Do(ctx)
	if err != nil {
		return 0, 0, err
	}
	return res.OrderID, res.TransactTime, nil
}

func (s *BinanceWrapper) GetOrder(ctx context.Context, symbol string, id int64) error {
	order, err := s.client.NewGetOrderService().
		Symbol(symbol).
		OrderID(id).
		Do(ctx)
	if err != nil {
		return err
	}
	d, _ := json.Marshal(order)
	logrus.Infof("order: %s", d)
	return nil
}

func (s *BinanceWrapper) GetSymbols(ctx context.Context) ([]binance.Symbol, error) {
	res, err := s.client.NewExchangeInfoService().Do(ctx)
	if err != nil {
		return nil, err
	}
	return res.Symbols, nil
}
