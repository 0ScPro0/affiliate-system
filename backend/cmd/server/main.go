package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/0ScPro0/affiliate-system/internal/core/config"
	core_postgres_pool "github.com/0ScPro0/affiliate-system/internal/core/database/postgres/pool"
	"github.com/0ScPro0/affiliate-system/internal/core/logger"
	core_http_middleware "github.com/0ScPro0/affiliate-system/internal/core/transport/http/middleware"
	core_http_server "github.com/0ScPro0/affiliate-system/internal/core/transport/http/server"
	city_postgres_repository "github.com/0ScPro0/affiliate-system/internal/features/affiliate/city/repository/postgres"
	city_service "github.com/0ScPro0/affiliate-system/internal/features/affiliate/city/service"
	city_transport_http "github.com/0ScPro0/affiliate-system/internal/features/affiliate/city/transport/http"
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

	fmt.Println(cfg.Database.DBUrl)

	log, err := logger.NewLogger(cfg)
	if err != nil {
		panic(err)
	}
	defer log.Close()
	log.Info("Logger successfully initialized")

	// Database pool
	pool, err := core_postgres_pool.NewConnectionPool(ctx, *cfg)
	if err != nil {
		panic(err)
	}
	defer pool.Close()
	log.Info("Database pool successfully initialized")

	// City feature
	cityRepository := city_postgres_repository.NewCityRepository(pool)
	cityService := city_service.NewCityService(cityRepository)
	cityHandler := city_transport_http.NewCityHTTPHandler(cityService)

	// Routes
	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegisterRoutes(cityHandler.Routes()...)

	httpServer := core_http_server.NewHTTPServer(
		cfg,
		log,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(log),
		core_http_middleware.Panic(),
		core_http_middleware.Trace(),
		core_http_middleware.Auth(cfg),
	)
	httpServer.RegisterAPIVersionRouter(*apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		log.Error("HTTP server run error:", zap.Error(err))
	}
}