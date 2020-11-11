package routes

import (
	"net/http"
	"regexp"
)

type xrout struct {
	pattern *regexp.Regexp
	handler http.Handler
}

// RegexpHandler - custom handler for routing via Regexp
type RegexpHandler struct {
	routes []*xrout
}

// Handler ...
func (rh *RegexpHandler) Handler(pattern *regexp.Regexp, handler http.Handler) {
	rh.routes = append(rh.routes, &xrout{pattern, handler})
}

// HandleFunc ...
func (rh *RegexpHandler) HandleFunc(pattern *regexp.Regexp, fhandler func(http.ResponseWriter, *http.Request)) {
	// Handler(pattern, http.HandlerFunc(fhandler))
	rh.routes = append(rh.routes, &xrout{pattern, http.HandlerFunc(fhandler)})
}

// HandleRegexp ...
func (rh *RegexpHandler) HandleRegexp(pattern string, fhandler func(http.ResponseWriter, *http.Request)) {
	// Handler(pattern, http.HandlerFunc(fhandler))
	rx := regexp.MustCompile(pattern)
	rh.routes = append(rh.routes, &xrout{rx, http.HandlerFunc(fhandler)})
}

// ServeHTTP ...
func (rh RegexpHandler) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	for _, rout := range rh.routes {
		if rout.pattern.MatchString(rq.URL.Path) {
			rout.handler.ServeHTTP(rw, rq)
			return
		}
	}
}
