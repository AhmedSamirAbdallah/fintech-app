package enum

type Currency string

const (
	USD = "USD"
	EGP = "EGP"
)

func IsValidCurrency(currency Currency) bool {
	switch currency {
	case USD, EGP:
		return true
	}
	return false
}
