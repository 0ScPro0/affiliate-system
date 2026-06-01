package main

import (
	"context"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/0ScPro0/affiliate-system/internal/core/config"
	"github.com/0ScPro0/affiliate-system/internal/core/logger"
	core_http_middleware "github.com/0ScPro0/affiliate-system/internal/core/transport/http/middleware"
	core_http_server "github.com/0ScPro0/affiliate-system/internal/core/transport/http/server"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	log, err := logger.NewLogger(cfg)
	if err != nil {
		panic(err)
	}
	defer log.Close()
	log.Info("Logger successfully initialized")

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	//apiVersionRouter.RegisterRoutes(...)

	httpServer := core_http_server.NewHTTPServer(
		cfg, 
		log,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(log),
		core_http_middleware.Panic(),
		core_http_middleware.Trace(),
	)
	httpServer.RegisterAPIVersionRouter(*apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		log.Error("HTTP server run error:", zap.Error(err))
	}
}