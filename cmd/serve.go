package cmd

import (
	"crypto-exchange-service/pkg/api"
	"crypto-exchange-service/pkg/config"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the http server",
	Run: func(cmd *cobra.Command, args []string) {
		// start the http server
		serve()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}

func serve() {
	var wg sync.WaitGroup

	settings := config.New()
	// start http servers
	s := api.NewServer(settings)
	s.Start(&wg)

	sig := make(chan os.Signal, 1024)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	for o := range sig {
		logrus.Infof("receive signal: %v", o)

		start := time.Now()
		// stop the server
		s.Stop()

		// wait for goroutines
		wg.Wait()

		logrus.Info("server is stopped")

		logrus.Infof("shut down takes time: %v", time.Since(start))
		return
	}
}
