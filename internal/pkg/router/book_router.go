package router

import (
	"github.com/go-chi/chi"
	"github.com/tanveerprottoy/book-service/internal/app/bookservice/module/book"
	"github.com/tanveerprottoy/book-service/internal/pkg/constant"
)

func RegisterBookRoutes(router *Router, version string, module *book.Module) {
	router.Mux.Group(
		func(r chi.Router) {
			r.Route(
				constant.ApiPattern+version+constant.BooksPattern,
				func(r chi.Router) {
					r.Get(constant.RootPattern, module.Handler.ReadMany)
					r.Get(constant.RootPattern+"{id}", module.Handler.ReadOne)
					r.Post(constant.RootPattern, module.Handler.Create)
					r.Patch(constant.RootPattern+"{id}", module.Handler.Update)
					r.Delete(constant.RootPattern+"{id}", module.Handler.Delete)
				},
			)
		},
	)
}
