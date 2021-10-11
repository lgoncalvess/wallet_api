package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	AuthenticationRequired bool
}

func Configure(r *mux.Router) *mux.Router {
	var routers []Router
	routers = append(routers, IO...)
	for _, router := range routers {
		r.HandleFunc(router.URI, router.Function).Methods(router.Method)
	}
	return r
}
