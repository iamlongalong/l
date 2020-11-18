package l

import (
	"net/http"
)

type LHandler struct {
	routes map[string]func(*Lres)
	req    *http.Request
	res    http.ResponseWriter
}

type Lres struct {
	res http.ResponseWriter
}
type Lroute = map[string]func(*Lres)

func NewHandler() *LHandler {
	routes := make(map[string]func(*Lres), 1)
	return &LHandler{
		routes: routes,
	}
}

func (r *Lres) Write(i interface{}) {
	switch i.(type) {
	case string:
		s := i.(string)
		r.res.Write([]byte(s))
	default:
		r.res.Write([]byte("233333"))
	}
}

func (h *LHandler) LoadRoutes(routes *Lroute) {
	for path, hand := range *routes {
		h.routes[path] = hand
	}
}

func (h *LHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if _, ok := h.routes[req.URL.Path]; !ok {
		res.WriteHeader(http.StatusNotFound)
	} else {
		lRes := &Lres{res}
		h.routes[req.URL.Path](lRes)
	}
}
