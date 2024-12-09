package enums

type UserRole int

const (
	USER_ROLE_PENDING = 0
	USER_ROLE_MEMBER  = 1
	USER_ROLE_ADMIN   = 2
)

func (enum UserRole) ToString() string {
	switch enum {
	case 1:
		return "member"
	case 2:
		return "admin"
	default:
		return "pending"
	}
}
