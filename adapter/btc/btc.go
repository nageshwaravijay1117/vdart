package btc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"vdart/config"
	"vdart/dtos"
)

//BTC --
type BTC struct {
	url string
}

//NewBTC --
func NewBTC() IBTC {
	return &BTC{
		url: config.BTCBaseURL,
	}
}

type IBTC interface {
	GetCurrencyDetailsBySymbol(symbol string) (dtos.SymbolDetail, error)
	GetAllCurrencies() ([]dtos.CurrencyDetail, error)
}

//GetCurrencyDetailsBySymbol -- Fetch Details of Currency from the BTC Endpoint
func (j *BTC) GetCurrencyDetailsBySymbol(symbol string) (dtos.SymbolDetail, error) {

	var resp dtos.SymbolDetail

	url := config.BTCBaseURL + config.GetAllCurrencyBySymbolURL + symbol
	fmt.Println("Currency URL is", url)
	response, err := http.Get(url)
	if err != nil {
		return resp, err
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Unable to read GetCurrencyDetailsBySymbol response body, Err:", err)
		return resp, err
	}

	fmt.Println("response", content)
	err = json.Unmarshal(content, &resp)
	if err != nil {
		fmt.Println("Unable to unmarshal GetCurrencyDetailsBySymbol response, Err:", err)
		return resp, err
	}
	return resp, nil

}

//GetAllCurrencies -- Fetch All Currency details from the BTC Endpoint
func (j *BTC) GetAllCurrencies() ([]dtos.CurrencyDetail, error) {

	resp := []dtos.CurrencyDetail{}

	url := config.BTCBaseURL + config.GetAllCurrenciesURL
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Unable to read GetAllCurrencies response body, Err:", err)
		return nil, err
	}

	err = json.Unmarshal(content, &resp)
	if err != nil {
		fmt.Println("Unable to unmarshal GetAllCurrencies response, Err:", err)
		return nil, err
	}
	return resp, nil

}
