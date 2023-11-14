package http

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/http/pprof"
	"sync"
)

type RouteConfig struct {
	Path    string
	Handler http.Handler
	Methods []string
}

type Router interface {
	AddRoute(config RouteConfig)
	StaticRoute()
	Mux() http.Handler
}

type router struct {
	mux *mux.Router
	sync.Mutex
}

func NewRouter() Router {
	r := &router{
		mux: mux.NewRouter().StrictSlash(true),
	}
	r.debugRoutes()
	r.healthCheckRoute()
	return r
}

func (r *router) AddRoute(route RouteConfig) {
	r.Lock()
	defer r.Unlock()
	handler := route.Handler

	rt := r.mux.Handle(route.Path, handler)
	if route.Methods != nil && len(route.Methods) > 0 {
		rt.Methods(route.Methods...)
	}

}

func (r *router) StaticRoute() {
	mux := r.mux
	mux.PathPrefix("/").Handler(http.FileServer(http.Dir("/tmp/static")))
}

func (r *router) Mux() http.Handler {
	return r.mux
}

func (r *router) debugRoutes() {
	mux := r.mux
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
}

func (r *router) healthCheckRoute() {
	r.mux.HandleFunc("/health_check", healthCheck)
}

func healthCheck(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Health Check Passed"))
}
