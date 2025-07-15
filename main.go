package main

import (
	"fmt"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/sawdustofmind/eth-balance-proxy/internal/config"
	"github.com/sawdustofmind/eth-balance-proxy/internal/log"
	"github.com/sawdustofmind/eth-balance-proxy/internal/monitoring"
	"github.com/sawdustofmind/eth-balance-proxy/internal/server"
	"github.com/sawdustofmind/eth-balance-proxy/internal/service"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal("loading config", zap.Error(err))
	}

	go func() {
		err := monitoring.RunMonitoringServer(cfg.MetricsPort)
		if err != nil {
			log.Error("Error starting monitoring server", zap.Error(err))
			return
		}
	}()

	// Gin server for eth balance
	bg, err := service.NewBalanceGetter(cfg.BalanceGetter)
	if err != nil {
		log.Fatal("creating balance getter", zap.Error(err))
	}

	r := gin.Default()
	r.Use(gin.Recovery(), monitoring.PrometheusMiddleware())

	impl := server.NewServerImpl(bg)
	server.RegisterHandlers(r, impl)

	monitoring.SetReady()

	r.Run(fmt.Sprintf(":%d", cfg.Port))
}
