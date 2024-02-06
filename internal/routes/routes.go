package routes

import (
	"dynamodb-crud/internal/repository/adapter"

	ServerConfig "dynamodb-crud/config/config"
	HealthHandler "dynamodb-crud/internal/handlers/health"
	ProductHandler "dynamodb-crud/internal/handlers/product"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Router struct {
	config *Config
	router *chi.Mux
}

func NewRouter() *Router {
	return &Router{
		config: NewConfig().SetTimeout(ServerConfig.GetConfig().Timeout),
		router: chi.NewRouter(),
	}
}

func (r *Router) SetRouters(repository adapter.Interface) *chi.Mux {
	r.setConfigRouters()
	r.RouterHealth(repository)
	r.RouterProduct(repository)

	return r.router
}

func (r *Router) setConfigRouters() {
	r.EnableCors()
	r.EnableLogger()
	r.EnableTimeout()
	r.EnableRecovery()
	r.EnableRequestID()
	r.EnableRealIP()
}

func (r *Router) RouterHealth(repository adapter.Interface) {
	handler := HealthHandler.NewHandler(repository)

	r.router.Route("/health", func(router chi.Router) {
		router.Post("/", handler.Post)
		router.Get("/", handler.Get)
		router.Put("/", handler.Put)
		router.Delete("/", handler.Delete)
		router.Options("/", handler.Options)
	})
}

func (r *Router) RouterProduct(repository adapter.Interface) {
	handler := ProductHandler.NewHandler(repository)

	r.router.Route("/product", func(router chi.Router) {
		router.Post("/", handler.Post)
		router.Get("/", handler.Get)
		router.Put("/{ID}", handler.Put)
		router.Delete("/{ID}", handler.Delete)
		router.Options("/", handler.Options)
	})
}

func (r *Router) EnableLogger() *Router {
	r.router.Use(middleware.Logger)
	return r
}

func (r *Router) EnableTimeout() *Router {
	r.router.Use(middleware.Timeout(r.config.GetTimeout()))
	return r
}

func (r *Router) EnableCors() *Router {
	r.router.Use(r.config.Cors)
	return r
}

func (r *Router) EnableRecovery() *Router {
	r.router.Use(middleware.Recoverer)
	return r
}

func (r *Router) EnableRequestID() *Router {
	r.router.Use(middleware.RequestID)
	return r
}

func (r *Router) EnableRealIP() *Router {
	r.router.Use(middleware.RealIP)
	return r
}
