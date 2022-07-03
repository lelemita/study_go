// 출처: tucker의 Go로 만드는 웹: https://youtu.be/YfrAlQKWRGg

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/lelemita/tuckerGoWeb/01_decorator_handler/decohandler"
	"github.com/lelemita/tuckerGoWeb/01_decorator_handler/myapp"
)

func main() {
	mux := NewHandler()
	http.ListenAndServe(":8080", mux)
}

func NewHandler() http.Handler {
	mux := myapp.NewHandler()
	logHandler := decohandler.NewHandler(mux, logFunc)
	logHandler2 := decohandler.NewHandler(logHandler, logFunc2)
	return logHandler2
}

func logFunc(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Println("[LOGGER1] Started")
	h.ServeHTTP(w, r)
	log.Println("[LOGGER1] Completed:", time.Since(start))
}

func logFunc2(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Println("[LOGGER2] Started")
	h.ServeHTTP(w, r)
	log.Println("[LOGGER2] Completed:", time.Since(start))
}
