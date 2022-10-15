package UT0311L04

import (
	"net"
	"reflect"
	"testing"
	"time"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-core/types"
	"github.com/uhppoted/uhppote-simulator/entities"
)

func TestNewUT0311L04(t *testing.T) {
	expected := UT0311L04{
		SerialNumber: types.SerialNumber(405060708),
		IpAddress:    net.IPv4(0, 0, 0, 0),
		SubnetMask:   net.IPv4(255, 255, 255, 0),
		Gateway:      net.IPv4(0, 0, 0, 0),
		// MacAddress:   types.MacAddress(mac),
		Version:  0x0892,
		Released: DefaultReleaseDate(),
		Doors: map[uint8]*entities.Door{
			1: entities.NewDoor(1),
			2: entities.NewDoor(2),
			3: entities.NewDoor(3),
			4: entities.NewDoor(4),
		},

		TimeProfiles: entities.TimeProfiles{},
		TaskList:     entities.TaskList{},
		Events:       entities.NewEventList(),
	}

	controller := NewUT0311L04(405060708, ".", false)

	if controller.SerialNumber != expected.SerialNumber {
		t.Errorf("incorrect serial number - expected:%v, got:%v", expected.SerialNumber, controller.SerialNumber)
	}

	if !reflect.DeepEqual(controller.IpAddress, expected.IpAddress) {
		t.Errorf("incorrect IP address - expected:%v, got:%v", expected.IpAddress, controller.IpAddress)
	}

	if !reflect.DeepEqual(controller.SubnetMask, expected.SubnetMask) {
		t.Errorf("incorrect netmask - expected:%v, got:%v", expected.SubnetMask, controller.SubnetMask)
	}

	if controller.Version != expected.Version {
		t.Errorf("incorrect version - expected:%v, got:%v", expected.Version, controller.Version)
	}

	if !reflect.DeepEqual(controller.Released, expected.Released) {
		t.Errorf("incorrect firmware date - expected:%v, got:%v", expected.Released, controller.Released)
	}

	if !reflect.DeepEqual(controller.Doors, expected.Doors) {
		t.Errorf("incorrect doors map - expected:%v, got:%v", expected.Doors, controller.Doors)
	}

	if !reflect.DeepEqual(controller.TimeProfiles, expected.TimeProfiles) {
		t.Errorf("incorrect time profiles list - expected:%v, got:%v", expected.TimeProfiles, controller.TimeProfiles)
	}

	if !reflect.DeepEqual(controller.TaskList, expected.TaskList) {
		t.Errorf("incorrect task list - expected:%v, got:%v", expected.TaskList, controller.TaskList)
	}

	if !reflect.DeepEqual(controller.Events, expected.Events) {
		t.Errorf("incorrect events list - expected:%v, got:%v", expected.Events, controller.Events)
	}
}

// TODO: ignore date/time fields
// func TestHandleGetStatus(t *testing.T) {
// 	swipeDateTime, _ := types.DateTimeFromString("2019-08-01 12:34:56")
// 	request := messages.GetStatusRequest{
// 		SerialNumber: 12345,
// 	}
//
// 	response := messages.GetStatusResponse{
// 		SerialNumber:  12345,
// 		EventIndex:     3,
// 		SwipeRecord:   0x00,
// 		Granted:       false,
// 		Door:          3,
// 		DoorOpened:    false,
// 		UserId:        1234567890,
// 		SwipeDateTime: *swipeDateTime,
// 		SwipeReason:   0x05,
// 		Door1State:    false,
// 		Door2State:    false,
// 		Door3State:    false,
// 		Door4State:    false,
// 		Door1Button:   false,
// 		Door2Button:   false,
// 		Door3Button:   false,
// 		Door4Button:   false,
// 		SystemState:   0x00,
// 		//	SystemDate     types.SystemDate   `uhppote:"offset:51"`
// 		//	SystemTime     types.SystemTime   `uhppote:"offset:37"`
// 		PacketNumber:   0,
// 		Backup:         0,
// 		SpecialMessage: 0,
// 		Battery:        0,
// 		FireAlarm:      0,
// 	}
//
// 	testHandle(&request, &response, t)
// }

func TestHandleOpenDoor(t *testing.T) {
	request := messages.OpenDoorRequest{
		SerialNumber: 12345,
		Door:         3,
	}

	response := messages.OpenDoorResponse{
		SerialNumber: 12345,
		Succeeded:    true,
	}

	testHandle(&request, &response, t)
}

func TestHandlePutCardRequest(t *testing.T) {
	from, _ := types.DateFromString("2019-01-01")
	to, _ := types.DateFromString("2019-12-31")
	request := messages.PutCardRequest{
		SerialNumber: 12345,
		CardNumber:   192837465,
		From:         from,
		To:           to,
		Door1:        1,
		Door2:        0,
		Door3:        1,
		Door4:        0,
	}

	response := messages.PutCardResponse{
		SerialNumber: 12345,
		Succeeded:    true,
	}

	testHandle(&request, &response, t)
}

func TestHandleDeleteCardRequest(t *testing.T) {
	request := messages.DeleteCardRequest{
		SerialNumber: 12345,
		CardNumber:   192837465,
	}

	response := messages.DeleteCardResponse{
		SerialNumber: 12345,
		Succeeded:    true,
	}

	testHandle(&request, &response, t)
}

func TestHandleDeleteCardsRequest(t *testing.T) {
	request := messages.DeleteCardsRequest{
		SerialNumber: 12345,
		MagicWord:    0x55aaaa55,
	}

	response := messages.DeleteCardsResponse{
		SerialNumber: 12345,
		Succeeded:    true,
	}

	testHandle(&request, &response, t)
}

func TestHandleGetCardsRequest(t *testing.T) {
	request := messages.GetCardsRequest{
		SerialNumber: 12345,
	}

	response := messages.GetCardsResponse{
		SerialNumber: 12345,
		Records:      3,
	}

	testHandle(&request, &response, t)
}

func TestHandleGetCardById(t *testing.T) {
	from, _ := types.DateFromString("2019-01-01")
	to, _ := types.DateFromString("2019-12-31")

	request := messages.GetCardByIDRequest{
		SerialNumber: 12345,
		CardNumber:   192837465,
	}

	response := messages.GetCardByIDResponse{
		SerialNumber: 12345,
		CardNumber:   192837465,
		From:         &from,
		To:           &to,
		Door1:        1,
		Door2:        0,
		Door3:        0,
		Door4:        1,
	}

	testHandle(&request, &response, t)
}

func TestHandleGetCardByIndex(t *testing.T) {
	from, _ := types.DateFromString("2019-01-01")
	to, _ := types.DateFromString("2019-12-31")

	request := messages.GetCardByIndexRequest{
		SerialNumber: 12345,
		Index:        2,
	}

	response := messages.GetCardByIndexResponse{
		SerialNumber: 12345,
		CardNumber:   192837465,
		From:         &from,
		To:           &to,
		Door1:        1,
		Door2:        0,
		Door3:        0,
		Door4:        1,
	}

	testHandle(&request, &response, t)
}

func TestHandleSetDoorControlState(t *testing.T) {
	request := messages.SetDoorControlStateRequest{
		SerialNumber: 12345,
		Door:         2,
		ControlState: 3,
		Delay:        7,
	}

	response := messages.SetDoorControlStateResponse{
		SerialNumber: 12345,
		Door:         2,
		ControlState: 3,
		Delay:        7,
	}

	testHandle(&request, &response, t)
}

func TestHandleGetDoorControlState(t *testing.T) {
	request := messages.GetDoorControlStateRequest{
		SerialNumber: 12345,
		Door:         2,
	}

	response := messages.GetDoorControlStateResponse{
		SerialNumber: 12345,
		Door:         2,
		ControlState: 2,
		Delay:        22,
	}

	testHandle(&request, &response, t)
}

func TestHandleSetListener(t *testing.T) {
	request := messages.SetListenerRequest{
		SerialNumber: 12345,
		Address:      net.IPv4(10, 0, 0, 1),
		Port:         43210,
	}

	response := messages.SetListenerResponse{
		SerialNumber: 12345,
		Succeeded:    true,
	}

	testHandle(&request, &response, t)
}

func TestHandleGetListener(t *testing.T) {
	request := messages.GetListenerRequest{
		SerialNumber: 12345,
	}

	response := messages.GetListenerResponse{
		SerialNumber: 12345,
		Address:      net.IPv4(10, 0, 0, 10),
		Port:         43210,
	}

	testHandle(&request, &response, t)
}

// TODO: deferred pending some way to compare Date field
// func TestHandleFindDevices(t *testing.T) {
// 	MAC, _ := net.ParseMAC("00:66:19:39:55:2d")
// 	now := types.Date(time.Now().UTC())
//
// 	request := messages.FindDevicesRequest{}
//
// 	response := messages.FindDevicesResponse{
// 		SerialNumber: 12345,
// 		IpAddress:    net.IPv4(10, 0, 0, 100),
// 		SubnetMask:   net.IPv4(255, 255, 255, 0),
// 		Gateway:      net.IPv4(10, 0, 0, 1),
// 		MacAddress:   types.MacAddress(MAC),
// 		Version:      9876,
// 		Date:         now,
// 	}
//
// 	testHandle(&request, &response, t)
// }

func TestHandleSetAddress(t *testing.T) {
	request := messages.SetAddressRequest{
		SerialNumber: 12345,
		Address:      net.IPv4(10, 0, 0, 100),
		Mask:         net.IPv4(255, 255, 255, 0),
		Gateway:      net.IPv4(10, 0, 0, 1),
		MagicWord:    0x55aaaa55,
	}

	testHandle(&request, nil, t)
}

func TestHandleGetEvent(t *testing.T) {
	datetime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2019-08-01 12:34:56", time.Local)
	timestamp := types.DateTime(datetime)

	request := messages.GetEventRequest{
		SerialNumber: 12345,
		Index:        2,
	}

	response := messages.GetEventResponse{
		SerialNumber: 12345,
		Index:        2,
		Type:         0x06,
		Granted:      true,
		Door:         4,
		Direction:    0x01,
		CardNumber:   555444321,
		Timestamp:    &timestamp,
		Reason:       9,
	}

	testHandle(&request, &response, t)
}

func TestHandleSetEventIndex(t *testing.T) {
	request := messages.SetEventIndexRequest{
		SerialNumber: 12345,
		Index:        7,
		MagicWord:    0x55aaaa55,
	}

	response := messages.SetEventIndexResponse{
		SerialNumber: 12345,
		Changed:      true,
	}

	testHandle(&request, &response, t)
}

func TestHandleGetEventIndex(t *testing.T) {
	request := messages.GetEventIndexRequest{
		SerialNumber: 12345,
	}

	response := messages.GetEventIndexResponse{
		SerialNumber: 12345,
		Index:        123,
	}

	testHandle(&request, &response, t)
}

func testHandle(request messages.Request, expected messages.Response, t *testing.T) {
	MAC, _ := net.ParseMAC("00:66:19:39:55:2d")
	from, _ := types.DateFromString("2019-01-01")
	to, _ := types.DateFromString("2019-12-31")
	timestamp := types.DateTime(time.Date(2019, time.August, 1, 12, 34, 56, 0, time.Local))
	listener := net.UDPAddr{IP: net.IPv4(10, 0, 0, 10), Port: 43210}

	doors := map[uint8]*entities.Door{
		1: &entities.Door{ControlState: 3, Delay: entities.DelayFromSeconds(11)},
		2: &entities.Door{ControlState: 2, Delay: entities.DelayFromSeconds(22)},
		3: &entities.Door{ControlState: 3, Delay: entities.DelayFromSeconds(33)},
		4: &entities.Door{ControlState: 3, Delay: entities.DelayFromSeconds(44)},
	}

	cards := entities.CardList{
		&entities.Card{
			CardNumber: 100000001,
			From:       &from,
			To:         &to,
			Doors:      map[uint8]uint8{1: 0, 2: 0, 3: 0, 4: 0},
		},
		&entities.Card{
			CardNumber: 192837465,
			From:       &from,
			To:         &to,
			Doors:      map[uint8]uint8{1: 1, 2: 0, 3: 0, 4: 1},
		},
		&entities.Card{
			CardNumber: 100000003,
			From:       &from,
			To:         &to,
			Doors:      map[uint8]uint8{1: 0, 2: 0, 3: 0, 4: 0},
		},
	}

	events := entities.MakeEventList(
		123,
		[]entities.Event{
			entities.Event{
				Index:     1,
				Type:      0x05,
				Granted:   false,
				Door:      3,
				Direction: 0x01,
				Card:      1234567890,
				Timestamp: &timestamp,
				Reason:    1,
			},
			entities.Event{
				Index:     2,
				Type:      0x06,
				Granted:   true,
				Door:      4,
				Direction: 0x01,
				Card:      555444321,
				Timestamp: &timestamp,
				Reason:    9,
			},
			entities.Event{
				Index:     3,
				Type:      0x05,
				Granted:   false,
				Door:      3,
				Direction: 0x01,
				Card:      1234567890,
				Timestamp: &timestamp,
				Reason:    1,
			},
			entities.Event{
				Index:     4,
				Type:      0x05,
				Granted:   true,
				Door:      3,
				Direction: 0x01,
				Card:      1234567891,
				Timestamp: &timestamp,
				Reason:    1,
			},
			entities.Event{
				Index:     5,
				Type:      0x05,
				Granted:   false,
				Door:      4,
				Direction: 0x01,
				Card:      1234567892,
				Timestamp: &timestamp,
				Reason:    1,
			},
			entities.Event{
				Index:     6,
				Type:      0x05,
				Granted:   false,
				Door:      1,
				Direction: 0x01,
				Card:      1234567892,
				Timestamp: &timestamp,
				Reason:    1,
			},
			entities.Event{
				Index:     7,
				Type:      0x05,
				Granted:   true,
				Door:      2,
				Direction: 0x02,
				Card:      1234567893,
				Timestamp: &timestamp,
				Reason:    1,
			},
			entities.Event{
				Index:     8,
				Type:      0x05,
				Granted:   true,
				Door:      3,
				Direction: 0x02,
				Card:      1234567894,
				Timestamp: &timestamp,
				Reason:    1,
			},
		})

	txq := make(chan entities.Message, 8)
	src := net.UDPAddr{IP: net.IPv4(10, 0, 0, 1), Port: 12345}

	s := UT0311L04{
		SerialNumber: 12345,
		IpAddress:    net.IPv4(10, 0, 0, 100),
		SubnetMask:   net.IPv4(255, 255, 255, 0),
		Gateway:      net.IPv4(10, 0, 0, 1),
		MacAddress:   types.MacAddress(MAC),
		Version:      9876,
		Listener:     &listener,
		Cards:        cards,
		Events:       events,
		Doors:        doors,

		txq: txq,
	}

	s.Handle(&src, request)

	if expected != nil {
		timeout := make(chan bool, 1)
		go func() {
			time.Sleep(100 * time.Millisecond)
			timeout <- true
		}()

		select {
		case response := <-txq:
			if response.Message == nil {
				t.Errorf("Invalid response: Expected: %v, got: %v", expected, response)
				return
			}

			if !reflect.DeepEqual(response.Message, expected) {
				t.Errorf("Incorrect response: Expected:\n%v, got:s\n%v", expected, response.Message)
			}

		case <-timeout:
			t.Errorf("No response from simulator")
		}
	}
}
