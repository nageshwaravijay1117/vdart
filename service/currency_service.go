package service

import (
	"errors"
	"fmt"
	"vdart/adapter/btc"

	"vdart/dtos"
)

type ICurrencyService interface {
	GetCurrencyDetailsBySymbol(symbol string) (dtos.SymbolDetail, error)
	GetAllCurrencies() (dtos.AllCurrencies, error)
}

type CurrencyService struct {
	btc btc.IBTC
}

// New returns a new service object
func NewCurrencyService() ICurrencyService {
	return &CurrencyService{
		btc: btc.NewBTC(),
	}
}

func (cs *CurrencyService) GetCurrencyDetailsBySymbol(symbol string) (dtos.SymbolDetail, error) {
	res, err := cs.btc.GetCurrencyDetailsBySymbol(symbol)
	if err != nil {
		fmt.Println("Error While Getting Data From Symbol")
	}
	if res.ID == "" {
		return res, errors.New("The Symbol Is Not Available")
	}
	return res, err
}

func (cs *CurrencyService) GetAllCurrencies() (dtos.AllCurrencies, error) {
	var response dtos.AllCurrencies
	res, err := cs.btc.GetAllCurrencies()
	if err != nil {
		fmt.Println("error while GetAllCurrencies data")
		return response, err
	}
	response.Currencies = res
	return response, nil
}
