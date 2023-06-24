package bookservice

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/tanveerprottoy/book-service/internal/app/bookservice/module/book"
	"github.com/tanveerprottoy/book-service/internal/pkg/constant"
	"github.com/tanveerprottoy/book-service/internal/pkg/router"
	"github.com/tanveerprottoy/book-service/pkg/data/sqlxpkg"
)

// App struct
type App struct {
	DBClient   *sqlxpkg.Client
	router     *router.Router
	Validate   *validator.Validate
	BookModule *book.Module
}

// Creates New App instance
func NewApp() *App {
	a := new(App)
	a.initComponents()
	return a
}

func (a *App) initDB() {
	a.DBClient = sqlxpkg.GetInstance()
}

func (a *App) initModuleRouters() {
	router.RegisterBookRoutes(a.router, constant.V1, a.BookModule)
}

func (a *App) initModules() {
	a.BookModule = book.NewModule(a.DBClient.DB, a.Validate)
}

// Init app
func (a *App) initComponents() {
	a.initDB()
	a.router = router.NewRouter()
	a.Validate = validator.New()
	a.initModules()
	a.initModuleRouters()
}

// Run app
func (a *App) Run() {
	err := http.ListenAndServe(
		":8080",
		a.router.Mux,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) RunDisableHTTP2() {
	srv := &http.Server{
		Handler:      a.router.Mux,
		Addr:         "127.0.0.1:8080",
		// TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}
	log.Fatal(srv.ListenAndServe())
}
