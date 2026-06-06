package main

import (
	"context"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/0ScPro0/affiliate-system/internal/core/config"
	core_postgres_pool "github.com/0ScPro0/affiliate-system/internal/core/database/postgres/pool"
	"github.com/0ScPro0/affiliate-system/internal/core/logger"
	core_http_middleware "github.com/0ScPro0/affiliate-system/internal/core/transport/http/middleware"
	core_http_server "github.com/0ScPro0/affiliate-system/internal/core/transport/http/server"
	auth_service "github.com/0ScPro0/affiliate-system/internal/features/affiliate/auth/service"
	auth_transport_http "github.com/0ScPro0/affiliate-system/internal/features/affiliate/auth/transport/http"
	category_postgres_repository "github.com/0ScPro0/affiliate-system/internal/features/affiliate/category/repository/postgres"
	category_service "github.com/0ScPro0/affiliate-system/internal/features/affiliate/category/service"
	category_transport_http "github.com/0ScPro0/affiliate-system/internal/features/affiliate/category/transport/http"
	city_postgres_repository "github.com/0ScPro0/affiliate-system/internal/features/affiliate/city/repository/postgres"
	city_service "github.com/0ScPro0/affiliate-system/internal/features/affiliate/city/service"
	city_transport_http "github.com/0ScPro0/affiliate-system/internal/features/affiliate/city/transport/http"
	partner_postgres_repository "github.com/0ScPro0/affiliate-system/internal/features/affiliate/partner/repository/postgres"
	partner_service "github.com/0ScPro0/affiliate-system/internal/features/affiliate/partner/service"
	partner_transport_http "github.com/0ScPro0/affiliate-system/internal/features/affiliate/partner/transport/http"
	user_postgres_repository "github.com/0ScPro0/affiliate-system/internal/features/affiliate/user/repository/postgres"
	user_service "github.com/0ScPro0/affiliate-system/internal/features/affiliate/user/service"
	user_transport_http "github.com/0ScPro0/affiliate-system/internal/features/affiliate/user/transport/http"

	_ "github.com/0ScPro0/affiliate-system/docs"
)

// @title           Affiliate System API
// @version         1.0
// @description     Affiliate System API
// @host            127.0.0.1:8080
// @BasePath        /api/v1/
// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization
// @description               Type "Bearer " followed by a space and your access token.
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

	// Database pool
	pool, err := core_postgres_pool.NewConnectionPool(ctx, *cfg)
	if err != nil {
		panic(err)
	}
	defer pool.Close()
	log.Info("Database pool successfully initialized")

	// Repositories
	userRepository := user_postgres_repository.NewUserRepository(pool)
	cityRepository := city_postgres_repository.NewCityRepository(pool)
	partnerRepository := partner_postgres_repository.NewPartnerRepository(pool)
	categoryRepository := category_postgres_repository.NewCategoryRepository(pool)

	// Services
	cityService := city_service.NewCityService(cityRepository)
	authService := auth_service.NewAuthService(cfg, userRepository)
	partnerService := partner_service.NewPartnerService(partnerRepository)
	categoryService := category_service.NewCategoryService(categoryRepository)
	userService := user_service.NewUserService(userRepository)

	// HTTP Handlers
	cityHandler := city_transport_http.NewCityHTTPHandler(cityService)
	authHandler := auth_transport_http.NewAuthHTTPHandler(authService)
	partnerHandler := partner_transport_http.NewPartnerHTTPHandler(partnerService)
	categoryHandler := category_transport_http.NewCategoryHTTPHandler(categoryService)
	userHandler := user_transport_http.NewUserHTTPHandler(userService)

	// Routes
	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegisterRoutes(cityHandler.Routes()...)
	apiVersionRouter.RegisterRoutes(authHandler.Routes()...)
	apiVersionRouter.RegisterRoutes(partnerHandler.Routes()...)
	apiVersionRouter.RegisterRoutes(categoryHandler.Routes()...)
	apiVersionRouter.RegisterRoutes(userHandler.Routes()...)

	httpServer := core_http_server.NewHTTPServer(
		cfg,
		log,
		core_http_middleware.CORS(),
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(log),
		core_http_middleware.Panic(),
		core_http_middleware.Trace(),
		core_http_middleware.Auth(cfg),
	)
	httpServer.RegisterAPIVersionRouter(*apiVersionRouter)
	httpServer.RegisterSwagger()

	if err := httpServer.Run(ctx); err != nil {
		log.Error("HTTP server run error:", zap.Error(err))
	}
}