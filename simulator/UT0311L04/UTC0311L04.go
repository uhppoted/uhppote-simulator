package UT0311L04

import (
	"bytes"
	"compress/gzip"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"reflect"
	"time"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-core/types"
	"github.com/uhppoted/uhppote-simulator/entities"
)

type UT0311L04 struct {
	file       string
	compressed bool
	touched    time.Time
	txq        chan entities.Message

	SerialNumber        types.SerialNumber    `json:"serial-number"`
	IpAddress           net.IP                `json:"address"`
	SubnetMask          net.IP                `json:"subnet"`
	Gateway             net.IP                `json:"gateway"`
	MacAddress          types.MacAddress      `json:"MAC"`
	Version             types.Version         `json:"version"`
	Released            *ReleaseDate          `json:"released"`
	TimeOffset          entities.Offset       `json:"offset"`
	Doors               entities.Doors        `json:"doors"`
	Listener            *net.UDPAddr          `json:"listener"`
	RecordSpecialEvents bool                  `json:"record-special-events"`
	PCControl           bool                  `json:"pc-control"`
	Interlock           uint8                 `json:"interlock"`
	SystemError         uint8                 `json:"system-error"`
	SequenceId          uint32                `json:"sequence-id"`
	SpecialInfo         uint8                 `json:"special-info"`
	InputState          uint8                 `json:"input-state"`
	TimeProfiles        entities.TimeProfiles `json:"time-profiles,omitempty"`
	TaskList            entities.TaskList     `json:"tasklist,omitempty"`
	Cards               entities.CardList     `json:"cards"`
	Events              entities.EventList    `json:"events"`
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
		Released:     DefaultReleaseDate(),
		Doors:        entities.MakeDoors(),
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

func (s *UT0311L04) SetTxQ(txq chan entities.Message) {
	s.txq = txq
}

func (s *UT0311L04) Handle(src *net.UDPAddr, rq messages.Request) {
	switch v := rq.(type) {
	case *messages.GetStatusRequest:
		s.getStatus(src, v)

	case *messages.SetTimeRequest:
		s.setTime(src, v)

	case *messages.GetTimeRequest:
		s.getTime(src, v)

	case *messages.OpenDoorRequest:
		s.unlockDoor(src, v)

	case *messages.PutCardRequest:
		s.putCard(src, v)

	case *messages.DeleteCardRequest:
		s.deleteCard(src, v)

	case *messages.DeleteCardsRequest:
		s.deleteCards(src, v)

	case *messages.GetCardsRequest:
		s.getCards(src, v)

	case *messages.GetCardByIDRequest:
		s.getCardByID(src, v)

	case *messages.GetCardByIndexRequest:
		s.getCardByIndex(src, v)

	case *messages.SetDoorControlStateRequest:
		s.setDoorControlState(src, v)

	case *messages.GetDoorControlStateRequest:
		s.getDoorControlState(src, v)

	case *messages.SetListenerRequest:
		s.setListener(src, v)

	case *messages.GetListenerRequest:
		s.getListener(src, v)

	case *messages.GetDeviceRequest:
		s.getDevice(src, v)

	case *messages.SetAddressRequest:
		s.setAddress(src, v)

	case *messages.GetEventRequest:
		s.getEvent(src, v)

	case *messages.SetEventIndexRequest:
		s.setEventIndex(src, v)

	case *messages.RecordSpecialEventsRequest:
		s.recordSpecialEvents(src, v)

	case *messages.GetEventIndexRequest:
		s.getEventIndex(src, v)

	case *messages.SetTimeProfileRequest:
		s.setTimeProfile(src, v)

	case *messages.GetTimeProfileRequest:
		s.getTimeProfile(src, v)

	case *messages.ClearTimeProfilesRequest:
		s.clearTimeProfiles(src, v)

	case *messages.ClearTaskListRequest:
		s.clearTaskList(src, v)

	case *messages.AddTaskRequest:
		s.addTask(src, v)

	case *messages.RefreshTaskListRequest:
		s.refreshTaskList(src, v)

	case *messages.SetPCControlRequest:
		s.setPCControl(src, v)

	case *messages.SetInterlockRequest:
		s.setInterlock(src, v)

	default:
		panic(fmt.Errorf("unsupported message type %T", v))
	}

	s.touched = time.Now()
}

func (s *UT0311L04) RunTasks() {
	handler := func(door uint8, task types.TaskType) {
		switch task {
		case types.DoorControlled:
			s.Doors.OverrideState(door, entities.Controlled)

		case types.DoorNormallyOpen:
			s.Doors.OverrideState(door, entities.NormallyOpen)

		case types.DoorNormallyClosed:
			s.Doors.OverrideState(door, entities.NormallyClosed)

		case types.DisableTimeProfile:
			s.Doors.EnableProfile(door, false)

		case types.EnableTimeProfile:
			s.Doors.EnableProfile(door, true)

			//	case types.CardNoPassword:
			//	case types.CardInPassword:
			//	case types.CardInOutPassword:
			//	case types.EnableMoreCards:
			//	case types.DisableMoreCards:

		case types.TriggerOnce:
			s.Doors.Unlock(door, 0*time.Second)

		case types.DisablePushButton:
			s.Doors.EnableButton(door, false)

		case types.EnablePushButton:
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

	buffer, err := io.ReadAll(zr)
	if err != nil {
		return nil, err
	}

	simulator := new(UT0311L04)
	err = json.Unmarshal(buffer, simulator)
	if err != nil {
		return nil, err
	}

	simulator.file = filepath
	simulator.compressed = true
	simulator.touched = time.Now()

	return simulator, nil
}

func load(filepath string) (*UT0311L04, error) {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	simulator := UT0311L04{
		Released:     DefaultReleaseDate(),
		Doors:        entities.MakeDoors(),
		TimeProfiles: entities.TimeProfiles{},
	}

	err = json.Unmarshal(bytes, &simulator)
	if err != nil {
		return nil, err
	}

	simulator.file = filepath
	simulator.compressed = false
	simulator.touched = time.Now()

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

func (s *UT0311L04) send(dest *net.UDPAddr, message interface{}) {
	if s.txq == nil {
		panic(fmt.Sprintf("Device %d: missing TXQ", s.SerialNumber))
	}

	if s.txq != nil && dest != nil && message != nil && !reflect.ValueOf(message).IsNil() {
		s.txq <- entities.Message{
			Destination: dest,
			Message:     message,
		}
	}
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
		Timestamp:  *event.Timestamp,
		CardNumber: event.Card,
		Granted:    event.Granted,
		Door:       event.Door,
		Direction:  event.Direction,
	}

	if fmt.Sprintf("%v", s.Version) == "6.62" {
		e662 := messages.EventV6_62{
			Event: e,
		}

		s.send(s.Listener, &e662)
	} else {
		s.send(s.Listener, &e)
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
