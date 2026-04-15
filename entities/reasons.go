package entities

type Reason uint8

const (
	ReasonUnknown               Reason = 0x00
	ReasonSwipePass             Reason = 0x01
	ReasonPCControl             Reason = 0x05
	ReasonNoPrivilege           Reason = 0x06
	ReasonInvalidPIN            Reason = 0x07
	ReasonAntiPassback          Reason = 0x08
	ReasonFirstCard             Reason = 0x0a
	ReasonNormallyClosed        Reason = 0x0b
	ReasonInvalidTimezone       Reason = 0x0f
	ReasonNoPass                Reason = 0x12
	ReasonPushbuttonOk          Reason = 0x14
	ReasonPushbuttonDisabled    Reason = 0x1e
	ReasonSuperPasswordOpenDoor Reason = 0x19
	ReasonInterlock             Reason = 0x21

	ReasonControllerPowerOn Reason = 0x1c
	ReasonControllerReset   Reason = 0x1d
)
