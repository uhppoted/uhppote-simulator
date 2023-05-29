package entities

const (
	ReasonOk              uint8 = 0
	ReasonSwipePass       uint8 = 0x01
	ReasonPCControl       uint8 = 0x05
	ReasonNoPrivilege     uint8 = 0x06
	ReasonInvalidPIN      uint8 = 0x07
	ReasonNormallyClosed  uint8 = 0x0b
	ReasonInvalidTimezone uint8 = 0x0f
	ReasonNoPass          uint8 = 0x12
	ReasonInterlock       uint8 = 0x21
)
