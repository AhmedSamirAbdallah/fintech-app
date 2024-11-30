package enum

type TransactionType string

const (
	Withdrawal TransactionType = "Withdrawal"
	Deposit    TransactionType = "Deposit"
	Transfer   TransactionType = "Transfer"
	Refund     TransactionType = "Refund"
)
