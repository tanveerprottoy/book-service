package router

import (
	"github.com/tanveerprottoy/book-service/pkg/middleware"

	"github.com/go-chi/chi"
)

// Router struct
type Router struct {
	Mux *chi.Mux
}

func NewRouter() *Router {
	r := &Router{}
	r.Mux = chi.NewRouter()
	r.registerGlobalMiddlewares()
	return r
}

func (r *Router) registerGlobalMiddlewares() {
	r.Mux.Use(
		middleware.JSONContentTypeMiddleWare,
		middleware.CORSEnableMiddleWare,
	)
}
