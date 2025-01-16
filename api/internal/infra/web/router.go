package web

import (
	"database/sql"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/jalvess021/kartka/api/internal/infra/web/routes"
	"fmt"
	"reflect"
	"runtime"
)

func SetupRoutes(r *chi.Mux, db *sql.DB) {
	
	//rota inicial 
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/api", http.StatusMovedPermanently)
	})

	r.Route("/api", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("avaiable routes"))
		})
		//Subrotas de produtos
		r.Mount("/products", routes.SetupProductRoutes(db))
	})

	chi.Walk(r, func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		handlerName := "unknown"
		if handlerFunc, ok := handler.(http.HandlerFunc); ok {
			handlerName = runtime.FuncForPC(reflect.ValueOf(handlerFunc).Pointer()).Name()
		}
		fmt.Printf("[%s]: %s -> %s has %d middlewares.\n", method, route, handlerName, len(middlewares))
		return nil
	})
}
