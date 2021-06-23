package dtos

type SymbolDetail struct {
	ID                   string `json:"id"`
	BaseCurrency         string `json:"baseCurrency"`
	QuoteCurrency        string `json:"quoteCurrency"`
	QuantityIncrement    string `json:"quantityIncrement"`
	TickSize             string `json:"tickSize"`
	TakeLiquidityRate    string `json:"takeLiquidityRate"`
	ProvideLiquidityRate string `json:"provideLiquidityRate"`
	FeeCurrency          string `json:"feeCurrency"`
	MarginTrading        bool   `json:"marginTrading"`
	MaxInitialLeverage   string `json:"maxInitialLeverage"`
}

type CurrencyDetail struct {
	ID                 string `json:"id"`
	FullName           string `json:"fullName"`
	Crypto             bool   `json:"crypto"`
	PayInEnabled       bool   `json:"payinEnabled"`
	PayInPaymentID     bool   `json:"payinPaymentId"`
	PayInConfirmations int    `json:"payinConfirmations"`
	PayoutEnabled      bool   `json:"payoutEnabled"`
	PayoutIsPaymentID  bool   `json:"payoutIsPaymentId"`
	TransferEnabled    bool   `json:"transferEnabled"`
	Delisted           bool   `json:"delisted"`
	PayoutFee          string `json:"payoutFee"`
	PrecisionPayout    int    `json:"precisionPayout"`
	PrecisionTransfer  int    `json:"precisionTransfer"`
}
type AllCurrencies struct {
	Currencies []CurrencyDetail `json:"currencies"`
}
