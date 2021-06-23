package handler

import (
	"net/http"
	"runtime/debug"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	//httpSwagger "github.com/swaggo/http-swagger"
)

type TokenRes struct {
	Exp int `json:"exp"`
}

// New creates a router and registers all the routes for the
// service and returns it.
func New() http.Handler {
	router := httprouter.New()
	router.PanicHandler = PanicHandler

	currencyRouter(router)
	//router.Handler("GET", "/swagger/*Path", httpSwagger.WrapHandler)
	handler := cors.AllowAll().Handler(router)
	return handler
}

func PanicHandler(w http.ResponseWriter, r *http.Request, c interface{}) {
	debug.PrintStack()
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(c.(error).Error()))
}
