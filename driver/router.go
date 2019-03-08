package driver

import (
	"net/http"
)

// Router implementation
// Does not use a dynamic map to hope to be slightly more performant
type Router struct {
	GetFunc    func(path string, handler http.Handler)
	PostFunc   func(path string, handler http.Handler)
	DeleteFunc func(path string, handler http.Handler)
	http.HandlerFunc
}

func NewRouter(getFunc func(path string, handler http.Handler), postFunc func(path string, handler http.Handler), deleteFunc func(path string, handler http.Handler), handlerFunc http.HandlerFunc) *Router {
	return &Router{GetFunc: getFunc, PostFunc: postFunc, DeleteFunc: deleteFunc, HandlerFunc: handlerFunc}
}

func (r *Router) Get(path string, handler http.Handler) {
	r.GetFunc(path, handler)
}

func (r *Router) Post(path string, handler http.Handler) {
	r.PostFunc(path, handler)
}

func (r *Router) Delete(path string, handler http.Handler) {
	r.DeleteFunc(path, handler)
}
