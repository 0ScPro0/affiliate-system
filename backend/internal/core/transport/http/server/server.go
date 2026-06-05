package core_http_server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"

	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"

	"github.com/0ScPro0/affiliate-system/docs"
	"github.com/0ScPro0/affiliate-system/internal/core/config"
	"github.com/0ScPro0/affiliate-system/internal/core/logger"
	core_http_middleware "github.com/0ScPro0/affiliate-system/internal/core/transport/http/middleware"
)

type HTTPServer struct {
	mux    *http.ServeMux
	config *config.ServerConfig
	log    *logger.Logger

	middleware []core_http_middleware.Middleware
}

func NewHTTPServer(
	cfg *config.Config,
	log *logger.Logger,
	middleware ...core_http_middleware.Middleware,
) *HTTPServer {
	return &HTTPServer{
		mux: http.NewServeMux(),
		config: &cfg.Server,
		log: log,
		middleware: middleware,
	}
}

func (s *HTTPServer) RegisterAPIVersionRouter(routers ...APIVersionRouter){
	for _, router := range routers{
		prefix := "/api/" + string(router.apiVersion)

		s.mux.Handle(
			prefix + "/", 
			http.StripPrefix(prefix, router),
		)
	}
}

func (s *HTTPServer) RegisterSwagger() {
	s.mux.Handle(
		"/docs/",
		httpSwagger.Handler(
			httpSwagger.URL("/docs/doc.json"),
		),
	)

	s.mux.HandleFunc(
		"/docs/doc.json",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(docs.SwaggerInfo.ReadDoc()))
		},
	)
}

func (s *HTTPServer) Run(ctx context.Context) error {
	mux := core_http_middleware.ChainMiddleware(s.mux, s.middleware...)

	addr := net.JoinHostPort(s.config.Host, strconv.Itoa(s.config.Port))
	server := &http.Server{
		Addr: addr,
		Handler: mux,
	}

	ch := make(chan error, 1)

	go func() {
		defer close(ch)

		s.log.Warn("Start HTTP server", zap.String("addr", addr))

		err := server.ListenAndServe()

		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("Listen and server HTTP: %w", err)
		}
	case <-ctx.Done():
		s.log.Warn("Shutdown HTTP server...")

		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			s.config.ShutdownTimeout,
		)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()

			return fmt.Errorf("Shutdown HTTP server: %w", err)
		}
	}

	s.log.Warn("HTTP server stopped")
	return nil

}