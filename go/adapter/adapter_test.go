package adapter

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var hello = "Hello World!"

type HelloHandler struct{}

func (h *HelloHandler) Handle(ctx *CustomHTTPContext) {
	if _, err := io.WriteString(ctx.ResponseWriter, hello); err != nil {
		ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
	}
}

func TestAdapter(t *testing.T) {
	adaptee := &HelloHandler{}
	adapted := HandlerAdapter(adaptee)
	AssertHTTPRequestReturnsHello(t, adapted)
}

func TestNetHttpAdapter(t *testing.T) {
	adaptee := func(ctx *CustomHTTPContext) error {
		if _, err := io.WriteString(ctx.ResponseWriter, hello); err != nil {
			ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		}
		return nil
	}

	adapted := HandlerFuncAdapter(adaptee)
	AssertHTTPRequestReturnsHello(t, adapted)

}

func AssertHTTPRequestReturnsHello(t *testing.T, adapted http.Handler) {
	ts := httptest.NewServer(adapted)
	defer ts.Close()

	res, err := http.Get(ts.URL)
	assert.NoError(t, err)

	message, err := io.ReadAll(res.Body)
	assert.Equal(t, hello, string(message))
}
