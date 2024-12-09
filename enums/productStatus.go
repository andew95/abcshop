package enums

type ProductStatus int

const (
	PRODUCT_STATUS_INACTIVE = 0
	PRODUCT_STATUS_ACTIVE   = 1
)

func (v ProductStatus) ToString() string {
	switch v {
	case PRODUCT_STATUS_INACTIVE:
		return "inactive"
	case PRODUCT_STATUS_ACTIVE:
		return "active"
	default:
		return "undefine"
	}
}
