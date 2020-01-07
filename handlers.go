package handlers

import "net/http"

// RegisterHandlers shows how to register each of the handler types on a mux
func RegisterHandlers(mux *http.ServeMux) {
	// RawHandler
	mux.Handle("/raw", &RawHandler{"hello"})

	// RawHandlerFunc
	rawHandlerFunc := &RawHandlerFunc{"hello"}
	mux.HandleFunc("/rawfunc", rawHandlerFunc.handle)

	// RawHandlerClosure
	rawHandlerClosure := &RawHandlerClosure{"hello"}
	mux.Handle("/rawclosure", rawHandlerClosure.handler())

	// Closure
	mux.HandleFunc("/closure", HandlerClosure("hello"))
}

// RawHandler is a struct that conforms to the http.Handler interface
type RawHandler struct {
	response string
}

var _ http.Handler = (*RawHandler)(nil)

func (raw *RawHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(raw.response))
}

// RawHandlerFunc is a struct with a method 'handle' that conforms to http.HandlerFunc
type RawHandlerFunc struct {
	response string
}

func (raw *RawHandlerFunc) handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(raw.response))
}

// RawHandlerClosure is a struct with a method handler() that returns a function that conforms to http.HandlerFunc
type RawHandlerClosure struct {
	response string
}

func (raw *RawHandlerClosure) handler() http.HandlerFunc {
	resp := []byte(raw.response)
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write(resp)
	}
}

// HandlerClosure is a function that returns a closure that conforms to http.HandlerFunc
func HandlerClosure(response string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(response))
	}
}
