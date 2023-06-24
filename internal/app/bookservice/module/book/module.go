package book

import (
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/tanveerprottoy/book-service/internal/app/bookservice/module/book/entity"
	"github.com/tanveerprottoy/book-service/pkg/data/sqlxpkg"
)

type Module struct {
	Handler    *Handler
	Service    *Service
	Repository sqlxpkg.Repository[entity.Book]
}

func NewModule(db *sqlx.DB, validate *validator.Validate) *Module {
	m := new(Module)
	// init order is reversed of the field decleration
	// as the dependency is served this way
	m.Repository = NewRepository(db)
	m.Service = NewService(m.Repository)
	m.Handler = NewHandler(m.Service, validate)
	return m
}
