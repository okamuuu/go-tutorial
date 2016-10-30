package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"

	myhandler "go-tutorial/echo/handler"
)

func TestHandler(t *testing.T) {
	e := echo.New()
	req := new(http.Request)
	rec := httptest.NewRecorder()
	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
	c.SetPath("/")

	handleFunc := myhandler.MainPage()
	handleFunc(c)
	if rec.Body.String() != "Hello World" {
		t.Errorf("expected response: Hello World, got %s", rec.Body.String())
	}
}
