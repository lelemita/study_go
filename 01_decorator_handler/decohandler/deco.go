package decohandler

import "net/http"

type DecoHandlerFunc func(w http.ResponseWriter, r *http.Request, h http.Handler)

type DecoHandler struct {
	h  http.Handler
	fn DecoHandlerFunc
}

func (d *DecoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d.fn(w, r, d.h)
}

func NewHandler(h http.Handler, fn DecoHandlerFunc) http.Handler {
	return &DecoHandler{
		h:  h,
		fn: fn,
	}
}
