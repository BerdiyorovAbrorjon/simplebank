package util

const (
	USD = "USD"
	UZS = "UZS"
	EUR = "EUR"
)

// check input string
func IsSupportCurrency(currency string) bool {
	switch currency {
	case USD, UZS, EUR:
		return true
	}
	return false
}
