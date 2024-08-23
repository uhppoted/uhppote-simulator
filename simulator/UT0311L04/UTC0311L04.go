package UT0311L04

import (
	"bytes"
	"compress/gzip"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/netip"
	"os"
	"path/filepath"
	"time"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-core/types"
	"github.com/uhppoted/uhppote-simulator/entities"
	"github.com/uhppoted/uhppote-simulator/log"
)

type UT0311L04 struct {
	file       string
	compressed bool
	touched    time.Time

	SerialNumber        types.SerialNumber    `json:"serial-number"`
	IpAddress           net.IP                `json:"address"`
	SubnetMask          net.IP                `json:"subnet"`
	Gateway             net.IP                `json:"gateway"`
	MacAddress          types.MacAddress      `json:"MAC"`
	Version             types.Version         `json:"version"`
	Released            types.Date            `json:"released,omitempty"`
	TimeOffset          entities.Offset       `json:"offset"`
	Doors               entities.Doors        `json:"doors"`
	Keypads             entities.Keypads      `json:"keypads"`
	Listener            netip.AddrPort        `json:"listener"`
	RecordSpecialEvents bool                  `json:"record-special-events"`
	PCControl           bool                  `json:"pc-control"`
	SystemError         uint8                 `json:"system-error"`
	SequenceId          uint32                `json:"sequence-id"`
	SpecialInfo         uint8                 `json:"special-info"`
	InputState          uint8                 `json:"input-state"`
	TimeProfiles        entities.TimeProfiles `json:"time-profiles,omitempty"`
	TaskList            entities.TaskList     `json:"tasklist,omitempty"`
	Cards               entities.CardList     `json:"cards"`
	Events              entities.EventList    `json:"events"`
}

var RELEASE_DATE = types.MustParseDate("2020-01-01")

var onEvent = func(dest netip.AddrPort, event any) {
}

func SetOnEvent(handler func(dest netip.AddrPort, event any)) {
	if handler != nil {
		onEvent = handler
	}
}

func NewUT0311L04(deviceID uint32, dir string, compressed bool) *UT0311L04 {
	filename := fmt.Sprintf("%d.json", deviceID)
	if compressed {
		filename = fmt.Sprintf("%d.json.gz", deviceID)
	}

	mac := []byte{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc}
	if _, err := rand.Read(mac); err != nil {
		fmt.Printf("   ... using default MAC address (%v)\n", err)
	}

	device := UT0311L04{
		file:       filepath.Join(dir, filename),
		compressed: compressed,
		touched:    time.Now(),

		SerialNumber: types.SerialNumber(deviceID),
		IpAddress:    net.IPv4(0, 0, 0, 0),
		SubnetMask:   net.IPv4(255, 255, 255, 0),
		Gateway:      net.IPv4(0, 0, 0, 0),
		MacAddress:   types.MacAddress(mac),
		Version:      0x0892,
		Released:     RELEASE_DATE,
		Doors:        entities.MakeDoors(),
		Keypads:      entities.MakeKeypads(),
		TimeProfiles: entities.TimeProfiles{},
		TaskList:     entities.TaskList{},
		Events:       entities.NewEventList(),
	}

	return &device
}

func (s *UT0311L04) DeviceID() uint32 {
	return uint32(s.SerialNumber)
}

func (s *UT0311L04) DeviceType() string {
	return "UT0311-L04"
}

func (s *UT0311L04) FilePath() string {
	return s.file
}

func (s *UT0311L04) Handle(rq messages.Request) (any, error) {
	s.touched = time.Now()

	switch v := rq.(type) {
	case *messages.ActivateAccessKeypadsRequest:
		return s.activateKeypads(v)

	case *messages.AddTaskRequest:
		return s.addTask(v)

	case *messages.ClearTaskListRequest:
		return s.clearTaskList(v)

	case *messages.ClearTimeProfilesRequest:
		return s.clearTimeProfiles(v)

	case *messages.DeleteCardRequest:
		return s.deleteCard(v)

	case *messages.DeleteCardsRequest:
		return s.deleteCards(v)

	case *messages.GetCardByIDRequest:
		return s.getCardByID(v)

	case *messages.GetCardByIndexRequest:
		return s.getCardByIndex(v)

	case *messages.GetCardsRequest:
		return s.getCards(v)

	case *messages.GetDeviceRequest:
		return s.getDevice(v)

	case *messages.GetDoorControlStateRequest:
		return s.getDoorControlState(v)

	case *messages.GetEventRequest:
		return s.getEvent(v)

	case *messages.GetEventIndexRequest:
		return s.getEventIndex(v)

	case *messages.GetListenerRequest:
		return s.getListener(v)

	case *messages.GetStatusRequest:
		return s.getStatus(v)

	case *messages.GetTimeRequest:
		return s.getTime(v)

	case *messages.GetTimeProfileRequest:
		return s.getTimeProfile(v)

	case *messages.OpenDoorRequest:
		return s.unlockDoor(v)

	case *messages.PutCardRequest:
		return s.putCard(v)

	case *messages.RecordSpecialEventsRequest:
		return s.recordSpecialEvents(v)

	case *messages.RefreshTaskListRequest:
		return s.refreshTaskList(v)

	case *messages.RestoreDefaultParametersRequest:
		return s.restoreDefaultParameters(v)

	case *messages.SetAddressRequest:
		return s.setAddress(v)

	case *messages.SetDoorControlStateRequest:
		return s.setDoorControlState(v)

	case *messages.SetDoorPasscodesRequest:
		return s.setDoorPasscodes(v)

	case *messages.SetEventIndexRequest:
		return s.setEventIndex(v)

	case *messages.SetInterlockRequest:
		return s.setInterlock(v)

	case *messages.SetListenerRequest:
		return s.setListener(v)

	case *messages.SetPCControlRequest:
		return s.setPCControl(v)

	case *messages.SetTimeProfileRequest:
		return s.setTimeProfile(v)

	case *messages.SetTimeRequest:
		return s.setTime(v)

	default:
		panic(fmt.Errorf("unsupported message type %T", v))
	}
}

func (s *UT0311L04) RunTasks() {
	handler := func(door uint8, task types.TaskType) {
		switch task {
		case types.DoorControlled:
			log.Infof("%-10v  task:set door %v to controlled", s.SerialNumber, door)
			s.Doors.OverrideState(door, entities.Controlled)

		case types.DoorNormallyOpen:
			log.Infof("%-10v  task:set door %v to normally open", s.SerialNumber, door)
			s.Doors.OverrideState(door, entities.NormallyOpen)

		case types.DoorNormallyClosed:
			log.Infof("%-10v  task:set door %v to normally closed", s.SerialNumber, door)
			s.Doors.OverrideState(door, entities.NormallyClosed)

		case types.DisableTimeProfile:
			log.Infof("%-10v  task:disabled time profile for door %v", s.SerialNumber, door)
			s.Doors.EnableProfile(door, false)

		case types.EnableTimeProfile:
			log.Infof("%-10v  task:enabled time profile for door %v", s.SerialNumber, door)
			s.Doors.EnableProfile(door, true)

		case types.CardNoPassword:
			log.Infof("%-10v  task:enabled card + no password for door %v", s.SerialNumber, door)
			s.Keypads[door] = entities.KeypadNone

		case types.CardInPassword:
			log.Infof("%-10v  task:enabled card + IN password for door %v", s.SerialNumber, door)
			s.Keypads[door] = entities.KeypadIn

		case types.CardInOutPassword:
			log.Infof("%-10v  task:enabled card + IN/OUT password for door %v", s.SerialNumber, door)
			s.Keypads[door] = entities.KeypadBoth

			//	case types.EnableMoreCards:
			//	case types.DisableMoreCards:

		case types.TriggerOnce:
			log.Infof("%-10v  task:trigger once for door %v", s.SerialNumber, door)
			s.Doors.Unlock(door, 0*time.Second)

		case types.DisablePushButton:
			log.Infof("%-10v  task:disabled push button for door %v", s.SerialNumber, door)
			s.Doors.EnableButton(door, false)

		case types.EnablePushButton:
			log.Infof("%-10v  task:enabled push button for door %v", s.SerialNumber, door)
			s.Doors.EnableButton(door, true)
		}
	}

	s.TaskList.Run(handler)
}

func Load(filepath string, compressed bool) (*UT0311L04, error) {
	if compressed {
		return loadGZ(filepath)
	}

	return load(filepath)
}

func loadGZ(filepath string) (*UT0311L04, error) {
	b, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	zr, err := gzip.NewReader(bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(zr)
	if err != nil {
		return nil, err
	}

	return unmarshal(bytes, filepath, true)
}

func load(filepath string) (*UT0311L04, error) {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	return unmarshal(bytes, filepath, false)
}

func unmarshal(bytes []byte, filepath string, compressed bool) (*UT0311L04, error) {
	object := struct {
		SerialNumber        types.SerialNumber    `json:"serial-number"`
		IpAddress           net.IP                `json:"address"`
		SubnetMask          net.IP                `json:"subnet"`
		Gateway             net.IP                `json:"gateway"`
		MacAddress          types.MacAddress      `json:"MAC"`
		Version             types.Version         `json:"version"`
		Released            types.Date            `json:"released,omitempty"`
		TimeOffset          entities.Offset       `json:"offset"`
		Doors               entities.Doors        `json:"doors"`
		Keypads             entities.Keypads      `json:"keypads"`
		Listener            json.RawMessage       `json:"listener"`
		RecordSpecialEvents bool                  `json:"record-special-events"`
		PCControl           bool                  `json:"pc-control"`
		SystemError         uint8                 `json:"system-error"`
		SequenceId          uint32                `json:"sequence-id"`
		SpecialInfo         uint8                 `json:"special-info"`
		InputState          uint8                 `json:"input-state"`
		TimeProfiles        entities.TimeProfiles `json:"time-profiles,omitempty"`
		TaskList            json.RawMessage       `json:"tasklist,omitempty"`
		Cards               entities.CardList     `json:"cards"`
		Events              entities.EventList    `json:"events"`
	}{
		Released:     RELEASE_DATE,
		Doors:        entities.MakeDoors(),
		Keypads:      entities.MakeKeypads(),
		TimeProfiles: entities.TimeProfiles{},
	}

	if err := json.Unmarshal(bytes, &object); err != nil {
		return nil, err
	}

	// ... unmarshal event listener variants
	var listener = netip.AddrPort{}
	var addrPort netip.AddrPort
	var udpAddr net.UDPAddr

	if err := json.Unmarshal(object.Listener, &addrPort); err == nil {
		listener = addrPort
	} else if err := json.Unmarshal(object.Listener, &udpAddr); err == nil {
		listener = udpAddr.AddrPort()
	}

	// ... unmarshal tasklist
	tasklist := struct {
		Tasks []types.Task `json:"tasks"`
	}{
		Tasks: []types.Task{},
	}

	if len(object.TaskList) > 0 {
		if err := json.Unmarshal(object.TaskList, &tasklist); err != nil {
			warnf(object.SerialNumber, "error loading tasklist (%v)", err)
		}
	}

	// ... initialise simulator
	simulator := UT0311L04{
		file:       filepath,
		compressed: compressed,
		touched:    time.Now(),

		SerialNumber:        object.SerialNumber,
		IpAddress:           object.IpAddress,
		SubnetMask:          object.SubnetMask,
		Gateway:             object.Gateway,
		MacAddress:          object.MacAddress,
		Version:             object.Version,
		Released:            object.Released,
		TimeOffset:          object.TimeOffset,
		Doors:               object.Doors,
		Keypads:             object.Keypads,
		Listener:            listener,
		RecordSpecialEvents: object.RecordSpecialEvents,
		PCControl:           object.PCControl,
		SystemError:         object.SystemError,
		SequenceId:          object.SequenceId,
		SpecialInfo:         object.SpecialInfo,
		InputState:          object.InputState,
		TimeProfiles:        object.TimeProfiles,
		TaskList: entities.TaskList{
			Tasks: tasklist.Tasks,
		},
		Cards:  object.Cards,
		Events: object.Events,
	}

	return &simulator, nil
}

func (s *UT0311L04) Save() error {
	if s.file != "" {
		if s.compressed {
			return saveGZ(s.file, s)
		}

		return save(s.file, s)
	}

	return nil
}

func (s *UT0311L04) Delete() error {
	if s.file != "" {
		if err := os.Remove(s.file); err != nil {
			return err
		}

		if _, err := os.Stat(s.file); err != nil {
			if !os.IsNotExist(err) {
				return err
			}
		}
	}

	return nil
}

func saveGZ(filepath string, s *UT0311L04) error {
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}

	var buffer bytes.Buffer
	zw := gzip.NewWriter(&buffer)
	_, err = zw.Write(b)
	if err != nil {
		return err
	}

	if err = zw.Close(); err != nil {
		return err
	}

	return os.WriteFile(filepath, buffer.Bytes(), 0644)
}

func save(filepath string, s *UT0311L04) error {
	bytes, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filepath, bytes, 0644)
}

func (s *UT0311L04) add(event entities.Event) {
	index := s.Events.Add(event)
	s.Save()

	utc := time.Now().UTC()
	datetime := utc.Add(time.Duration(s.TimeOffset))

	e := messages.Event{
		SerialNumber: s.SerialNumber,
		EventIndex:   index,
		SystemError:  s.SystemError,
		SystemDate:   types.SystemDate(datetime),
		SystemTime:   types.SystemTime(datetime),
		SequenceId:   s.SequenceId,
		SpecialInfo:  s.SpecialInfo,
		RelayState:   s.relays(),
		InputState:   s.InputState,

		Door1State: s.Doors.IsOpen(1),
		Door2State: s.Doors.IsOpen(2),
		Door3State: s.Doors.IsOpen(3),
		Door4State: s.Doors.IsOpen(4),

		Door1Button: s.Doors.IsButtonPressed(1),
		Door2Button: s.Doors.IsButtonPressed(2),
		Door3Button: s.Doors.IsButtonPressed(3),
		Door4Button: s.Doors.IsButtonPressed(4),

		EventType:  event.Type,
		Reason:     event.Reason,
		Timestamp:  event.Timestamp,
		CardNumber: event.Card,
		Granted:    event.Granted,
		Door:       event.Door,
		Direction:  event.Direction,
	}

	// ... firmware 6.62 had a slightly different format
	if fmt.Sprintf("%v", s.Version) == "6.62" {
		e662 := messages.EventV6_62{
			Event: e,
		}

		onEvent(s.Listener, &e662)
	} else {
		onEvent(s.Listener, &e)
	}
}

func (s *UT0311L04) relays() uint8 {
	state := uint8(0x00)
	doors := map[uint8]uint8{
		1: 0x01,
		2: 0x02,
		3: 0x04,
		4: 0x08,
	}

	for k, mask := range doors {
		if s.Doors.IsUnlocked(k) {
			state |= mask
		}
	}

	return state
}

func warnf(tag any, format string, args ...any) {
	f := fmt.Sprintf("%-10v  %v", tag, format)

	log.Warnf(f, args...)
}
