package models

type OrderStatus uint32

const (
	OrderStatusOpen   OrderStatus = 0
	OrderStatusFilled OrderStatus = 1
	OrderStatusFailed OrderStatus = 2
)

type OrderV1 struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement:true"`
	UserAddress string `gorm:"index:idx_user"`
	InputToken  string
	Amount      string
	OutputToken string
	Status      uint32
}

func NewOrderV1(userAddress, inputToken, amount, outputToken string) *OrderV1 {
	return &OrderV1{
		Status:      uint32(OrderStatusOpen),
		UserAddress: userAddress,
		InputToken:  inputToken,
		Amount:      amount,
		OutputToken: outputToken,
	}
}
