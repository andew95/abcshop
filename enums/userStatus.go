package enums

type UserStatus int

const (
	USER_STATUS_INACTIVE = 0
	USER_STATUS_ACTIVE   = 1
	USER_STATUS_PENDING  = 2
)

func (v UserStatus) ToString() string {
	switch v {
	case USER_STATUS_INACTIVE:
		return "inactive"
	case USER_STATUS_ACTIVE:
		return "active"
	default:
		return "pending"
	}
}
