package routes

import (
	"net/http"
	"regexp"
)

// Hrout ...
func Hrout() {
	//fmt.Println("My rout mod")
}

type xrout struct {
	pattern *regexp.Regexp
	handler http.Handler
}

type regexpHandler struct {
	routes []*xrout
}

// Handler ...
func (rh *regexpHandler) Handler(pattern *regexp.Regexp, handler http.Handler) {
	rh.routes = append(rh.routes, &xrout{pattern, handler})
}

// HandleFunc ...
func (rh *regexpHandler) HandleFunc(pattern *regexp.Regexp, fhandler func(http.ResponseWriter, *http.Request)) {
	// Handler(pattern, http.HandlerFunc(fhandler))
	rh.routes = append(rh.routes, &xrout{pattern, http.HandlerFunc(fhandler)})
}

// HandleFunc ...
func (rh *regexpHandler) HandleRegexp(pattern string, fhandler func(http.ResponseWriter, *http.Request)) {
	// Handler(pattern, http.HandlerFunc(fhandler))
	rx := regexp.MustCompile(pattern)
	rh.routes = append(rh.routes, &xrout{rx, http.HandlerFunc(fhandler)})
}

// ServeHTTP ...
func (rh *regexpHandler) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	for _, rout := range rh.routes {
		if rout.pattern.MatchString(rq.URL.Path) {
			rout.handler.ServeHTTP(rw, rq)
			return
		}
	}
}
