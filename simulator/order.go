package simulator

type OrderType int32
type OrderID string
type OrderAction int32

const (
	OrderTypeUnknown OrderType = 0
	OrderTypeLimit   OrderType = 1
	OrderTypeMarket  OrderType = 2
)

func (t OrderType) String() string {
	switch t {
	case OrderTypeLimit:
		return "limit"

	case OrderTypeMarket:
		return "market"
	}

	return "unknown"
}

const (
	OrderActionUnknown           = 0
	OrderActionNewOrder          = 4
	OrderActionAcceptNewOrder    = 4 | 1
	OrderActionRejectNewOrder    = 4 | 2
	OrderActionCancelOrder       = 8
	OrderActionAcceptCancelOrder = 8 | 1
	OrderActionRejectCancelOrder = 8 | 2
)

func (a OrderAction) IsRequest() bool {
	return (a & 3) == 0
}

func (a OrderAction) IsResponse() bool {
	return (a & 3) != 0
}

func (a OrderAction) IsAccept() bool {
	return (a & 1) == 1
}

func (a OrderAction) IsReject() bool {
	return (a & 2) == 2
}

func (a OrderAction) String() string {
	// TODO: implement
	return "unknown"
}

type Order interface {
	Symbol() string
	Quantity() int64
	Price() float64
	Type() OrderType
	OrderID() OrderID
}
