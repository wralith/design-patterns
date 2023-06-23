package adapter

import (
	"net/http"
)

// You should really check the test code for this example since codes to be adapted are in there!
// CustomHTTPHandler -> http.Handler

// Struct level adapter example

type CustomHTTPHandler interface {
	Handle(c *CustomHTTPContext)
}

type CustomHTTPContext struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
	// ...
	// Bunch of other stuff maybe
}

// HandlerAdapter adapts  CustomHTTPHandler to http.Handler
func HandlerAdapter(handler CustomHTTPHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.Handle(&CustomHTTPContext{
			Request:        r,
			ResponseWriter: w,
		})
	})
}

// Function level example

type CustomHTTPHandlerFunc func(c *CustomHTTPContext) error

// HandlerFuncAdapter adapts CustomHTTPHandlerFunc to http.HandlerFunc
func HandlerFuncAdapter(handler CustomHTTPHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := &CustomHTTPContext{
			Request:        r,
			ResponseWriter: w,
		}
		err := handler(c)
		if err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
		}
	}
}
