package serverhttp

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"

	healthhndl "go-template/internal/transport/server/http/health"
	pkglog "go-template/pkg/log"
)

func NewRouter(log *zap.Logger) *chi.Mux {
	router := chi.NewRouter()

	router.Use(
		middleware.Recoverer,
		middleware.RequestLogger(pkglog.NewLogFormatter(log)),
	)

	router.Get("/health", healthhndl.NewHandler())

	return router
}
