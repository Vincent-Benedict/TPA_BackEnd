package api

import(
	"github.com/Vincent-Benedict/TPA-Web/middleware"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router{
	r:= mux.NewRouter()
	r.Use(middleware.LogMiddleware)

	return r
}
