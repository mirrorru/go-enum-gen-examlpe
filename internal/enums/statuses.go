package enums

//go:generate go-enum -f $GOFILE --marshal --sql --names --values --nocomments --output-suffix _enum

// ClientStatus - статус (состояние) клиента
// ENUM(new,active,inactive,blocked)
type ClientStatus byte

func (x ClientStatus) HumanString() string {
	switch x {
	case ClientStatusNew:
		return "Новый"
	case ClientStatusActive:
		return "Активный"
	case ClientStatusInactive:
		return "Неактивный"
	case ClientStatusBlocked:
		return "Заблокирован"
	default:
		return ""
	}
}
