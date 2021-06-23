package config

var (
	Port                      string
	GetAllCurrenciesURL       string
	GetAllCurrencyBySymbolURL string
	BTCBaseURL                string
)

func init() {
	Port = "3001"
	GetAllCurrencyBySymbolURL = "api/2/public/symbol/"
	GetAllCurrenciesURL = "api/2/public/currency"
	BTCBaseURL = "https://api.hitbtc.com/"
}
