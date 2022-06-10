package routing

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jacky-htg/go-services/controllers"
	"github.com/jmoiron/sqlx"
)

func API(db *sqlx.DB, log *log.Logger) http.Handler {
	app := newApp(log)
	v := controllers.Users{DB: db, Log: log}

	app.Handle(http.MethodGet, "/users", v.List)
	app.Handle(http.MethodGet, "/users/{id}", v.View)
	app.Handle(http.MethodPost, "/users", v.Create)
	app.Handle(http.MethodPut, "/users/{id}", v.Update)
	app.Handle(http.MethodDelete, "/users/{id}", v.Delete)

	return app
}

func newApp(log *log.Logger) *App {
	return &App{
		log: log,
		mux: chi.NewRouter(),
	}
}

type App struct {
	log *log.Logger
	mux *chi.Mux
}

func (a *App) Handle(method, url string, h http.Handler) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)

		if err != nil {
			a.log.Printf("ERROR: %+v", err)
			if err := api.Response(w, err); err != nil {
				a.log.Printf("ERROR: %+v", err)
			}
		}
	}
	a.mux.MethodFunc(method, url, h)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
