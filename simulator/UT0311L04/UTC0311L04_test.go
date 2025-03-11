package UT0311L04

import (
	"net"
	"net/netip"
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
		Version:      0x0892,
		Released:     RELEASE_DATE,
		Doors:        entities.MakeDoors(),
		TimeProfiles: entities.TimeProfiles{},
		TaskList:     entities.TaskList{},
		Events:       entities.NewEventList(),
		AntiPassback: entities.AntiPassback{},
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

	if !reflect.DeepEqual(controller.TaskList.Tasks, expected.TaskList.Tasks) {
		t.Errorf("incorrect task list - expected:%v, got:%v", expected.TaskList.Tasks, controller.TaskList.Tasks)
	}

	if !reflect.DeepEqual(controller.Events, expected.Events) {
		t.Errorf("incorrect events list - expected:%v, got:%v", expected.Events, controller.Events)
	}
}

func TestUT0311L04Unmarshal(t *testing.T) {
	MAC, _ := net.ParseMAC("00:12:23:34:45:56")
	doors := entities.MakeDoors()

	expected := UT0311L04{
		file:       "test.json",
		compressed: false,

		SerialNumber: types.SerialNumber(405419896),
		IpAddress:    net.IPv4(0, 0, 0, 0),
		SubnetMask:   net.IPv4(255, 255, 255, 0),
		Gateway:      net.IPv4(0, 0, 0, 0),
		MacAddress:   types.MacAddress(MAC),
		Version:      0x0892,
		Released:     types.MustParseDate("2018-11-05"),
		Listener:     netip.MustParseAddrPort("192.168.1.100:60001"),
		Doors:        doors,
		Keypads:      map[uint8]entities.Keypad{1: 3, 2: 3, 3: 0, 4: 3},
		TimeProfiles: entities.TimeProfiles{},
		TaskList: entities.TaskList{
			Tasks: []types.Task{
				types.Task{
					Task: types.DoorNormallyClosed,
					Door: 4,
					From: types.MustParseDate("2024-01-01"),
					To:   types.MustParseDate("2024-12-31"),
					Weekdays: types.Weekdays{
						time.Sunday:    false,
						time.Monday:    true,
						time.Tuesday:   false,
						time.Wednesday: false,
						time.Thursday:  false,
						time.Friday:    true,
						time.Saturday:  false,
					},
					Start: types.NewHHmm(8, 30),
					Cards: 0,
				},
			},
		},
		Events:       entities.NewEventList(),
		AntiPassback: entities.MakeAntiPassback(types.Readers1_234),
	}

	JSON := `
	{
	  "serial-number": 405419896,
	  "address": "0.0.0.0",
	  "subnet": "255.255.255.0",
	  "gateway": "0.0.0.0",
	  "MAC": "00:12:23:34:45:56",
	  "version": "0892",
	  "released": "2018-11-05",
	  "offset": "-7h0m0.666172s",
	  "doors": {
	    "1": { "control": 3, "delay": "5s", "passcodes": [ 0, 0, 0, 0 ] },
	    "2": { "control": 3, "delay": "5s", "passcodes": [ 0, 0, 0, 0 ] },
	    "3": { "control": 3, "delay": "5s", "passcodes": [ 0, 0, 0, 0 ] },
	    "4": { "control": 3, "delay": "5s", "passcodes": [ 0, 0, 0, 0 ] },
	    "interlock": 4
	  },
	  "keypads": { "1": 3, "2": 3, "3": 0, "4": 3 },
	  "listener": "192.168.1.100:60001",
	  "record-special-events": true,
	  "pc-control": true,
	  "system-error": 0,
	  "sequence-id": 0,
	  "special-info": 0,
	  "input-state": 0,
	  "time-profiles": {
	    "100": {
	      "id": 100,
	      "start-date": "2023-04-01",
	      "end-date": "2023-12-31",
	      "weekdays": "Monday,Wednesday,Friday",
	      "segments": [
	        { "start": "08:30", "end": "11:30" },
	        { "start": "00:00", "end": "00:00" },
	        { "start": "13:45", "end": "17:00" }
	      ]
	    }
	  },
	  "tasklist": {
	    "tasks": [
	      { "task": "LOCK DOOR", "door": 4, "start-date": "2024-01-01", "end-date": "2024-12-31", "weekdays": "Monday,Friday", "start": "08:30" }
	    ]
	  },
	  "cards": [],
	  "events": {
	    "size": 256,
	    "chunk": 8,
	    "index": 0,
	    "events": [
	      { "index": 1, "type": 1, "granted": false, "door": 3, "direction": 1, "card": 10058400, "timestamp": "2024-05-08 10:50:53", "reason": 18 }
	    ]
	  },
	  "anti-passback": 4
	}
	`

	if controller, err := unmarshal([]byte(JSON), "test.json", false); err != nil {
		t.Fatalf("error unmarshalling UT0311L04 (%v)", err)
	} else if controller == nil {
		t.Fatalf("error unmarshalling UT0311L04, expected UT0311L04{}, got %v", controller)
	} else {
		expected.touched = controller.touched

		if controller.SerialNumber != expected.SerialNumber {
			t.Errorf("incorrect serial number\n   expected:%v\n   got:    %v", expected.SerialNumber, controller.SerialNumber)
		}

		if !reflect.DeepEqual(controller.IpAddress, expected.IpAddress) {
			t.Errorf("incorrect IP address\n   expected:%v\n   got:    %v", expected.IpAddress, controller.IpAddress)
		}

		if !reflect.DeepEqual(controller.SubnetMask, expected.SubnetMask) {
			t.Errorf("incorrect netmask\n   expected:%v\n   got:    %v", expected.SubnetMask, controller.SubnetMask)
		}

		if controller.Version != expected.Version {
			t.Errorf("incorrect version\n   expected:%v\n   got:    %v", expected.Version, controller.Version)
		}

		if !reflect.DeepEqual(controller.Released, expected.Released) {
			t.Errorf("incorrect firmware date\n   expected:%v\n   got:    %v", expected.Released, controller.Released)
		}

		if !reflect.DeepEqual(controller.TaskList.Tasks, expected.TaskList.Tasks) {
			t.Errorf("incorrect tasklist\n   expected:%v\n   got:    %v", expected.TaskList.Tasks, controller.TaskList.Tasks)
		}

		if !reflect.DeepEqual(controller.AntiPassback, expected.AntiPassback) {
			t.Errorf("incorrect anti-passback\n   expected:%v\n   got:     %v", expected.AntiPassback, controller.AntiPassback)
		}
	}
}

func TestUT0311L04UnmarshalDefaultValues(t *testing.T) {
	expected := UT0311L04{
		file:       "test.json",
		compressed: false,
		Released:   RELEASE_DATE,
	}

	JSON := `
	{
	}
	`

	if controller, err := unmarshal([]byte(JSON), "test.json", false); err != nil {
		t.Fatalf("error unmarshalling UT0311L04 (%v)", err)
	} else if controller == nil {
		t.Fatalf("error unmarshalling UT0311L04, expected UT0311L04{}, got %v", controller)
	} else {
		expected.touched = controller.touched

		if !reflect.DeepEqual(controller.Released, expected.Released) {
			t.Errorf("incorrect firmware date - expected:%v, got:%v", expected.Released, controller.Released)
		}
	}
}

func TestUT0311L04UnmarshalWithMissingTaskList(t *testing.T) {
	MAC, _ := net.ParseMAC("00:12:23:34:45:56")
	doors := entities.MakeDoors()

	expected := UT0311L04{
		file:       "test.json",
		compressed: false,

		SerialNumber: types.SerialNumber(405419896),
		IpAddress:    net.IPv4(0, 0, 0, 0),
		SubnetMask:   net.IPv4(255, 255, 255, 0),
		Gateway:      net.IPv4(0, 0, 0, 0),
		MacAddress:   types.MacAddress(MAC),
		Version:      0x0892,
		Released:     types.MustParseDate("2018-11-05"),
		Listener:     netip.MustParseAddrPort("192.168.1.100:60001"),
		Doors:        doors,
		Keypads:      map[uint8]entities.Keypad{1: 3, 2: 3, 3: 0, 4: 3},
		TimeProfiles: entities.TimeProfiles{},
		TaskList: entities.TaskList{
			Tasks: []types.Task{},
		},
		Events: entities.NewEventList(),
	}

	JSON := `
	{
	  "serial-number": 405419896,
	  "address": "0.0.0.0",
	  "subnet": "255.255.255.0",
	  "gateway": "0.0.0.0",
	  "MAC": "00:12:23:34:45:56",
	  "version": "0892",
	  "released": "2018-11-05",
	  "offset": "-7h0m0.666172s",
	  "doors": {
	    "1": { "control": 3, "delay": "5s", "passcodes": [ 0, 0, 0, 0 ] },
	    "2": { "control": 3, "delay": "5s", "passcodes": [ 0, 0, 0, 0 ] },
	    "3": { "control": 3, "delay": "5s", "passcodes": [ 0, 0, 0, 0 ] },
	    "4": { "control": 3, "delay": "5s", "passcodes": [ 0, 0, 0, 0 ] },
	    "interlock": 4
	  },
	  "keypads": { "1": 3, "2": 3, "3": 0, "4": 3 },
	  "listener": "192.168.1.100:60001",
	  "record-special-events": true,
	  "pc-control": true,
	  "system-error": 0,
	  "sequence-id": 0,
	  "special-info": 0,
	  "input-state": 0,
	  "time-profiles": {
	    "100": {
	      "id": 100,
	      "start-date": "2023-04-01",
	      "end-date": "2023-12-31",
	      "weekdays": "Monday,Wednesday,Friday",
	      "segments": [
	        { "start": "08:30", "end": "11:30" },
	        { "start": "00:00", "end": "00:00" },
	        { "start": "13:45", "end": "17:00" }
	      ]
	    }
	  },
	  "cards": [],
	  "events": {
	    "size": 256,
	    "chunk": 8,
	    "index": 0,
	    "events": [
	      { "index": 1, "type": 1, "granted": false, "door": 3, "direction": 1, "card": 10058400, "timestamp": "2024-05-08 10:50:53", "reason": 18 }
	    ]
	  }
	}
	`

	if controller, err := unmarshal([]byte(JSON), "test.json", false); err != nil {
		t.Fatalf("error unmarshalling UT0311L04 (%v)", err)
	} else if controller == nil {
		t.Fatalf("error unmarshalling UT0311L04, expected UT0311L04{}, got %v", controller)
	} else {
		expected.touched = controller.touched

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

		if !reflect.DeepEqual(controller.TaskList.Tasks, expected.TaskList.Tasks) {
			t.Errorf("incorrect tasklist - expected:%v, got:%v", expected.TaskList.Tasks, controller.TaskList.Tasks)
		}
	}
}

func TestUT0311L04UnmarshalWithInvalidTaskList(t *testing.T) {
	MAC, _ := net.ParseMAC("00:12:23:34:45:56")
	doors := entities.MakeDoors()

	expected := UT0311L04{
		file:       "test.json",
		compressed: false,

		SerialNumber: types.SerialNumber(405419896),
		IpAddress:    net.IPv4(0, 0, 0, 0),
		SubnetMask:   net.IPv4(255, 255, 255, 0),
		Gateway:      net.IPv4(0, 0, 0, 0),
		MacAddress:   types.MacAddress(MAC),
		Version:      0x0892,
		Released:     types.MustParseDate("2018-11-05"),
		Listener:     netip.MustParseAddrPort("192.168.1.100:60001"),
		Doors:        doors,
		Keypads:      map[uint8]entities.Keypad{1: 3, 2: 3, 3: 0, 4: 3},
		TimeProfiles: entities.TimeProfiles{},
		TaskList: entities.TaskList{
			Tasks: []types.Task{},
		},
		Events: entities.NewEventList(),
	}

	JSON := `
	{
	  "serial-number": 405419896,
	  "address": "0.0.0.0",
	  "subnet": "255.255.255.0",
	  "gateway": "0.0.0.0",
	  "MAC": "00:12:23:34:45:56",
	  "version": "0892",
	  "released": "2018-11-05",
	  "offset": "-7h0m0.666172s",
	  "doors": {
	    "1": { "control": 3, "delay": "5s", "passcodes": [ 0, 0, 0, 0 ] },
	    "2": { "control": 3, "delay": "5s", "passcodes": [ 0, 0, 0, 0 ] },
	    "3": { "control": 3, "delay": "5s", "passcodes": [ 0, 0, 0, 0 ] },
	    "4": { "control": 3, "delay": "5s", "passcodes": [ 0, 0, 0, 0 ] },
	    "interlock": 4
	  },
	  "keypads": { "1": 3, "2": 3, "3": 0, "4": 3 },
	  "listener": "192.168.1.100:60001",
	  "record-special-events": true,
	  "pc-control": true,
	  "system-error": 0,
	  "sequence-id": 0,
	  "special-info": 0,
	  "input-state": 0,
	  "time-profiles": {
	    "100": {
	      "id": 100,
	      "start-date": "2023-04-01",
	      "end-date": "2023-12-31",
	      "weekdays": "Monday,Wednesday,Friday",
	      "segments": [
	        { "start": "08:30", "end": "11:30" },
	        { "start": "00:00", "end": "00:00" },
	        { "start": "13:45", "end": "17:00" }
	      ]
	    }
	  },
	  "tasklist": "notasksheremovealong",
	  "cards": [],
	  "events": {
	    "size": 256,
	    "chunk": 8,
	    "index": 0,
	    "events": [
	      { "index": 1, "type": 1, "granted": false, "door": 3, "direction": 1, "card": 10058400, "timestamp": "2024-05-08 10:50:53", "reason": 18 }
	    ]
	  }
	}
	`

	if controller, err := unmarshal([]byte(JSON), "test.json", false); err != nil {
		t.Fatalf("error unmarshalling UT0311L04 (%v)", err)
	} else if controller == nil {
		t.Fatalf("error unmarshalling UT0311L04, expected UT0311L04{}, got %v", controller)
	} else {
		expected.touched = controller.touched

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

		if !reflect.DeepEqual(controller.TaskList.Tasks, expected.TaskList.Tasks) {
			t.Errorf("incorrect tasklist - expected:%v, got:%v", expected.TaskList.Tasks, controller.TaskList.Tasks)
		}
	}
}

func testHandle(request messages.Request, expected messages.Response, t *testing.T) {
	MAC, _ := net.ParseMAC("00:66:19:39:55:2d")
	from := types.MustParseDate("2019-01-01")
	to := types.MustParseDate("2019-12-31")
	timestamp := types.DateTime(time.Date(2019, time.August, 1, 12, 34, 56, 0, time.Local))
	listener := netip.MustParseAddrPort("10.0.0.10:43210")
	doors := entities.MakeDoors()

	doors.SetControlState(1, 3)
	doors.SetControlState(2, 2)
	doors.SetControlState(3, 3)
	doors.SetControlState(4, 3)

	doors.SetDelay(1, entities.DelayFromSeconds(11))
	doors.SetDelay(2, entities.DelayFromSeconds(22))
	doors.SetDelay(3, entities.DelayFromSeconds(33))
	doors.SetDelay(4, entities.DelayFromSeconds(44))

	cards := entities.CardList{
		&entities.Card{
			CardNumber: 100000001,
			From:       from,
			To:         to,
			Doors:      map[uint8]uint8{1: 0, 2: 0, 3: 0, 4: 0},
		},
		&entities.Card{
			CardNumber: 192837465,
			From:       from,
			To:         to,
			Doors:      map[uint8]uint8{1: 1, 2: 0, 3: 0, 4: 1},
		},
		&entities.Card{
			CardNumber: 100000003,
			From:       from,
			To:         to,
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
				Timestamp: timestamp,
				Reason:    1,
			},
			entities.Event{
				Index:     2,
				Type:      0x06,
				Granted:   true,
				Door:      4,
				Direction: 0x01,
				Card:      555444321,
				Timestamp: timestamp,
				Reason:    9,
			},
			entities.Event{
				Index:     3,
				Type:      0x05,
				Granted:   false,
				Door:      3,
				Direction: 0x01,
				Card:      1234567890,
				Timestamp: timestamp,
				Reason:    1,
			},
			entities.Event{
				Index:     4,
				Type:      0x05,
				Granted:   true,
				Door:      3,
				Direction: 0x01,
				Card:      1234567891,
				Timestamp: timestamp,
				Reason:    1,
			},
			entities.Event{
				Index:     5,
				Type:      0x05,
				Granted:   false,
				Door:      4,
				Direction: 0x01,
				Card:      1234567892,
				Timestamp: timestamp,
				Reason:    1,
			},
			entities.Event{
				Index:     6,
				Type:      0x05,
				Granted:   false,
				Door:      1,
				Direction: 0x01,
				Card:      1234567892,
				Timestamp: timestamp,
				Reason:    1,
			},
			entities.Event{
				Index:     7,
				Type:      0x05,
				Granted:   true,
				Door:      2,
				Direction: 0x02,
				Card:      1234567893,
				Timestamp: timestamp,
				Reason:    1,
			},
			entities.Event{
				Index:     8,
				Type:      0x05,
				Granted:   true,
				Door:      3,
				Direction: 0x02,
				Card:      1234567894,
				Timestamp: timestamp,
				Reason:    1,
			},
		})

	s := UT0311L04{
		SerialNumber: 12345,
		IpAddress:    net.IPv4(10, 0, 0, 100),
		SubnetMask:   net.IPv4(255, 255, 255, 0),
		Gateway:      net.IPv4(10, 0, 0, 1),
		MacAddress:   types.MacAddress(MAC),
		Version:      9876,
		Listener:     listener,
		AntiPassback: entities.MakeAntiPassback(types.Readers1_234),
		Cards:        cards,
		Events:       events,
		Doors:        doors,
	}

	if response, err := s.Handle(request); err != nil {
		t.Fatalf("%v", err)
	} else if expected == nil && response != nil {
		t.Errorf("invalid response - expected:%v, got:%v", expected, response)
	} else if expected != nil && response == nil {
		t.Errorf("Invalid response: Expected: %v, got: %v", expected, response)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("Incorrect response\n   expected:%v\n   got:     %v", expected, response)
	}
}
