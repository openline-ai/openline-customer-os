package enum

type Currency string

const (
	CurrencyAUD Currency = "AUD"
	CurrencyBRL Currency = "BRL"
	CurrencyCAD Currency = "CAD"
	CurrencyCHF Currency = "CHF"
	CurrencyCNY Currency = "CNY"
	CurrencyEUR Currency = "EUR"
	CurrencyGBP Currency = "GBP"
	CurrencyHKD Currency = "HKD"
	CurrencyINR Currency = "INR"
	CurrencyJPY Currency = "JPY"
	CurrencyKRW Currency = "KRW"
	CurrencyMXN Currency = "MXN"
	CurrencyNOK Currency = "NOK"
	CurrencyNZD Currency = "NZD"
	CurrencyRON Currency = "RON"
	CurrencySEK Currency = "SEK"
	CurrencySGD Currency = "SGD"
	CurrencyTRY Currency = "TRY"
	CurrencyUSD Currency = "USD"
	CurrencyZAR Currency = "ZAR"
)

func (e Currency) String() string {
	return string(e)
}

func DecodeCurrency(code string) Currency {
	switch code {
	case "USD":
		return CurrencyUSD
	case "EUR":
		return CurrencyEUR
	case "GBP":
		return CurrencyGBP
	case "JPY":
		return CurrencyJPY
	case "AUD":
		return CurrencyAUD
	case "CAD":
		return CurrencyCAD
	case "CHF":
		return CurrencyCHF
	case "CNY":
		return CurrencyCNY
	case "SEK":
		return CurrencySEK
	case "NZD":
		return CurrencyNZD
	case "KRW":
		return CurrencyKRW
	case "SGD":
		return CurrencySGD
	case "NOK":
		return CurrencyNOK
	case "MXN":
		return CurrencyMXN
	case "INR":
		return CurrencyINR
	case "HKD":
		return CurrencyHKD
	case "BRL":
		return CurrencyBRL
	case "ZAR":
		return CurrencyZAR
	case "TRY":
		return CurrencyTRY
	case "RON":
		return CurrencyRON
	default:
		return ""
	}
}

func (e Currency) Symbol() string {
	switch e {
	case CurrencyUSD:
		return "$"
	case CurrencyEUR:
		return "€"
	case CurrencyGBP:
		return "£"
	case CurrencyJPY:
		return "¥"
	case CurrencyAUD:
		return "A$"
	case CurrencyCAD:
		return "C$"
	case CurrencyCHF:
		return "Fr"
	case CurrencyCNY:
		return "¥"
	case CurrencySEK:
		return "kr"
	case CurrencyNZD:
		return "NZ$"
	case CurrencyKRW:
		return "₩"
	case CurrencySGD:
		return "S$"
	case CurrencyNOK:
		return "kr"
	case CurrencyMXN:
		return "Mex$"
	case CurrencyINR:
		return "₹"
	case CurrencyHKD:
		return "HK$"
	case CurrencyBRL:
		return "R$"
	case CurrencyZAR:
		return "R"
	case CurrencyTRY:
		return "₺"
	case CurrencyRON:
		return "L"
	default:
		return ""
	}
}
