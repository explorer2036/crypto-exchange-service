package api

import (
	"net"
	"sync"
	"time"

	"crypto-exchange-service/pkg/api/exchange"
	"crypto-exchange-service/pkg/config"

	"github.com/buaazp/fasthttprouter"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

// Server for handling fasthttp requests
type Server struct {
	server         *fasthttp.Server
	settings       *config.Config
	validate       *validator.Validate
	binanceWrapper *exchange.BinanceWrapper
	done           chan struct{}
}

// NewServer returns a fasthttp server
func NewServer(settings *config.Config) *Server {
	s := &Server{
		settings: settings,
		validate: validator.New(),
		done:     make(chan struct{}),
	}

	router := fasthttprouter.New()
	router.POST("/order", s.Order)

	s.binanceWrapper = exchange.NewBinanceWrapper(settings)

	s.server = &fasthttp.Server{
		Handler:      router.Handler,
		ReadTimeout:  time.Duration(settings.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(settings.WriteTimeout) * time.Second,
	}

	return s
}

// Start the fasthttp server
func (s *Server) Start(wg *sync.WaitGroup) {
	wg.Add(1)

	listener, err := net.Listen("tcp", s.settings.Address)
	if err != nil {
		panic(err)
	}
	go func() {
		wg.Done()

		logrus.Infof("http server started on %s", s.settings.Address)
		if err := s.server.Serve(listener); err != nil {
			logrus.Errorf("http listen and serve: %v", err)
		}
	}()
}

// Stop the fasthttp server
func (s *Server) Stop() {
	if s.done != nil {
		close(s.done)
	}

	// shut down the fasthttp server
	if s.server != nil {
		if err := s.server.Shutdown(); err != nil {
			logrus.Errorf("shutdown fasthttp server: %v", err)
		}
	}
}
