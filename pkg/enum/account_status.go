package enum

type AccountStatus string

const (
	Active    = "Active"
	Inactive  = "Inactive"
	Suspended = "Suspended"
)

func IsValidAccountStatus(status AccountStatus) bool {
	switch status {
	case Active, Inactive, Suspended:
		return true
	}
	return false
}
