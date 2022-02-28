package cmd

import (
	"context"
	"crypto-exchange-service/apis/order"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	address string
)

func init() {
	OrderCmd.PersistentFlags().StringVar(&address, "address", "localhost:9001", "Crypto exchange service's address")
}

func init() {
	RootCmd.AddCommand(OrderCmd)
}

func newClient(addr string) (order.OrderServiceClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return order.NewOrderServiceClient(conn), nil
}

var OrderCmd = &cobra.Command{
	Use:   "order",
	Short: "Create a new order",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := newClient(address)
		if err != nil {
			logrus.Errorf("New client: %v", err)
			return
		}

		request := &order.CreateOrderRequest{
			Exchange: "binance",
			MarketId: "LTCUSDT",
			Side:     "buy",
			Quantity: 0.2,
			Price:    125,
		}
		response, err := client.CreateOrder(context.Background(), request)
		if err != nil {
			logrus.Errorf("Create order: %v", err)
			return
		}
		logrus.Infof("order id: %d, timestamp: %d", response.OrderId, response.Timestamp)
	},
}
