package handler

import (
	"errors"
	"net/http"

	"vdart/service"

	"github.com/julienschmidt/httprouter"
)

func currencyRouter(router *httprouter.Router) {
	router.GET("/currency/symbol/:symbol", GetCurrencyDetailsBySymbol)
	router.GET("/currency/all", GetAllCurrencies)
}

func GetCurrencyDetailsBySymbol(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rd := logAndGetContext(w, r)

	symbol := ps.ByName("symbol")

	if symbol == "" {
		writeJSONMessage(errors.New("Symbol Cannot be empty").Error(), TypeErrMsg, http.StatusInternalServerError, rd)
	}

	currencyService := service.NewCurrencyService()
	res, err := currencyService.GetCurrencyDetailsBySymbol(symbol)
	if err != nil {
		writeJSONMessage(err.Error(), TypeErrMsg, http.StatusInternalServerError, rd)
		return
	}
	writeJSONStruct(res, http.StatusOK, rd)

}

func GetAllCurrencies(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rd := logAndGetContext(w, r)

	currencyService := service.NewCurrencyService()
	res, err := currencyService.GetAllCurrencies()
	if err != nil {
		writeJSONMessage(err.Error(), TypeErrMsg, http.StatusInternalServerError, rd)
		return
	}
	writeJSONStruct(res, http.StatusOK, rd)

}
