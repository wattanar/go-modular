package router

import (
	"go-modular/internal/app/handler"

	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	r := httprouter.New()

	r.GET("/", handler.Home)

	return r
}
