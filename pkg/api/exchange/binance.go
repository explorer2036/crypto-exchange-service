package exchange

import (
	"context"

	"crypto-exchange-service/pkg/config"

	"github.com/adshao/go-binance/v2"
)

// BinanceWrapper represents the wrapper for the Binance exchange.
type BinanceWrapper struct {
	client *binance.Client
}

// NewBinanceWrapper creates a generic wrapper of the binance API.
func NewBinanceWrapper(settings *config.Config) *BinanceWrapper {
	if settings.TestNetwork == "yes" {
		binance.UseTestnet = true
	}
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

func (s *BinanceWrapper) GetOrder(ctx context.Context, symbol string, id int64) (*binance.Order, error) {
	order, err := s.client.NewGetOrderService().
		Symbol(symbol).
		OrderID(id).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *BinanceWrapper) CancelOrder(ctx context.Context, symbol string, id int64) (*binance.CancelOrderResponse, error) {
	res, err := s.client.NewCancelOrderService().
		Symbol(symbol).
		OrderID(id).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *BinanceWrapper) GetSymbols(ctx context.Context) ([]binance.Symbol, error) {
	res, err := s.client.NewExchangeInfoService().Do(ctx)
	if err != nil {
		return nil, err
	}
	return res.Symbols, nil
}
