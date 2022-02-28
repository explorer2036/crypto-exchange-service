package api

import (
	"net"
	"sync"
	"time"

	"crypto-exchange-service/apis/order"
	"crypto-exchange-service/pkg/api/exchange"
	"crypto-exchange-service/pkg/config"

	"github.com/buaazp/fasthttprouter"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/grpc"
)

// Server for handling fasthttp requests
type Server struct {
	httpServer     *fasthttp.Server
	grpcServer     *grpc.Server
	settings       *config.Config
	validate       *validator.Validate
	binanceWrapper *exchange.BinanceWrapper
}

// NewServer returns a fasthttp server
func NewServer(settings *config.Config) *Server {
	s := &Server{
		settings: settings,
		validate: validator.New(),
	}

	router := fasthttprouter.New()
	router.POST("/order", s.Order)

	s.binanceWrapper = exchange.NewBinanceWrapper(settings)

	s.httpServer = &fasthttp.Server{
		Handler:      router.Handler,
		ReadTimeout:  time.Duration(settings.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(settings.WriteTimeout) * time.Second,
	}

	s.grpcServer = grpc.NewServer()
	order.RegisterOrderServiceServer(s.grpcServer, NewOrderServiceImpl(s.binanceWrapper))

	return s
}

// Start the fasthttp server
func (s *Server) Start(wg *sync.WaitGroup) {
	httpListener, err := net.Listen("tcp", s.settings.HTTPAddr)
	if err != nil {
		panic(err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		logrus.Infof("http server started on %s", s.settings.HTTPAddr)
		if err := s.httpServer.Serve(httpListener); err != nil {
			logrus.Errorf("http listen and serve: %v", err)
		}
	}()

	grpcListener, err := net.Listen("tcp", s.settings.GRPCAddr)
	if err != nil {
		panic(err)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()

		logrus.Infof("grpc server started on %s", s.settings.GRPCAddr)
		if err := s.grpcServer.Serve(grpcListener); err != nil {
			logrus.Errorf("grpc listen and serve: %v", err)
		}
	}()
}

// Stop the fasthttp server
func (s *Server) Stop() {
	// shut down the fasthttp server
	if s.httpServer != nil {
		if err := s.httpServer.Shutdown(); err != nil {
			logrus.Errorf("shutdown fasthttp server: %v", err)
		}
	}
	// shut down the grpc server
	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
	}

}
