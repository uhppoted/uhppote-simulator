package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/uhppoted/uhppote-simulator/simulator"
)

type Device struct {
	DeviceID   uint32 `json:"device-id"`
	DeviceType string `json:"device-type"`
}

type DeviceList struct {
	Devices []Device `json:"devices"`
}

type NewDeviceRequest struct {
	DeviceID   uint32 `json:"device-id"`
	DeviceType string `json:"device-type"`
	Compressed bool   `json:"compressed"`
}

type handlerfn func(*simulator.Context, http.ResponseWriter, *http.Request)

type handler struct {
	re *regexp.Regexp
	fn handlerfn
}

type dispatcher struct {
	ctx      *simulator.Context
	handlers []handler
}

func Run(ctx *simulator.Context) {
	d := dispatcher{
		ctx,
		make([]handler, 0),
	}

	d.Add("^/uhppote/simulator$", devices)
	d.Add("^/uhppote/simulator/[0-9]+$", device)
	d.Add("^/uhppote/simulator/[0-9]+/swipe$", swipe)
	d.Add("^/uhppote/simulator/[0-9]+/door/[1-4]$", door)

	log.Fatal(http.ListenAndServe(ctx.RestAddress, &d))
}

func (d *dispatcher) Add(path string, h handlerfn) {
	re := regexp.MustCompile(path)
	d.handlers = append(d.handlers, handler{re, h})
}

func (d *dispatcher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// CORS header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	// CORS pre-flight request ?
	if r.Method == http.MethodOptions {
		return
	}

	url := r.URL.Path
	for _, h := range d.handlers {
		if h.re.MatchString(url) {
			h.fn(d.ctx, w, r)
			return
		}
	}

	http.Error(w, "Unsupported API", http.StatusBadRequest)
}

func devices(ctx *simulator.Context, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		list(ctx, w, r)

	case http.MethodPost:
		create(ctx, w, r)

	default:
		http.Error(w, fmt.Sprintf("Invalid method:%s - expected GET or POST", r.Method), http.StatusMethodNotAllowed)
	}
}

func device(ctx *simulator.Context, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		delete(ctx, w, r)

	default:
		http.Error(w, fmt.Sprintf("Invalid method:%s - expected DELETE", r.Method), http.StatusMethodNotAllowed)
	}
}

func list(ctx *simulator.Context, w http.ResponseWriter, r *http.Request) {
	_, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request", http.StatusInternalServerError)
		return
	}

	devices := make([]Device, 0)

	ctx.DeviceList.Apply(func(s simulator.Simulator) {
		devices = append(devices, Device{
			DeviceID:   s.DeviceID(),
			DeviceType: s.DeviceType(),
		})
	})

	response := DeviceList{devices}
	b, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error generating response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func create(ctx *simulator.Context, w http.ResponseWriter, r *http.Request) {
	blob, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request", http.StatusInternalServerError)
		return
	}

	request := NewDeviceRequest{}
	err = json.Unmarshal(blob, &request)
	if err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	if request.DeviceID < 1 {
		http.Error(w, "Missing device ID", http.StatusBadRequest)
		return
	}

	if request.DeviceType != "UT0311-L04" {
		http.Error(w, "Invalid  device type - expected UT0311-L04", http.StatusBadRequest)
		return
	}

	created, err := ctx.DeviceList.Add(request.DeviceID, request.Compressed, ctx.Directory)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating device %d: %v", request.DeviceID, err), http.StatusInternalServerError)
		return
	}

	if created {
		w.Header().Set("Location", fmt.Sprintf("/uhppote/simulator/%d", request.DeviceID))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func delete(ctx *simulator.Context, w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	matches := regexp.MustCompile("^/uhppote/simulator/([0-9]+)$").FindStringSubmatch(url)
	deviceID, err := strconv.ParseUint(matches[1], 10, 32)
	if err != nil {
		http.Error(w, "Error reading request", http.StatusInternalServerError)
		return
	}

	_, err = ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request", http.StatusInternalServerError)
		return
	}

	if s := ctx.DeviceList.Find(uint32(deviceID)); s == nil {
		http.Error(w, fmt.Sprintf("No device with ID %d", deviceID), http.StatusNotFound)
		return
	}

	ctx.DeviceList.Delete(uint32(deviceID))

	w.Header().Set("Content-Type", "application/json")
}

func swipe(ctx *simulator.Context, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, fmt.Sprintf("Invalid method:%s - expected POST", r.Method), http.StatusMethodNotAllowed)
		return
	}

	url := r.URL.Path
	matches := regexp.MustCompile("^/uhppote/simulator/([0-9]+)/swipe$").FindStringSubmatch(url)
	deviceID, err := strconv.ParseUint(matches[1], 10, 32)
	if err != nil {
		http.Error(w, "Error reading request", http.StatusInternalServerError)
		return
	}

	blob, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request", http.StatusInternalServerError)
		return
	}

	request := struct {
		Door       uint8  `json:"door"`
		CardNumber uint32 `json:"card-number"`
	}{}

	err = json.Unmarshal(blob, &request)
	if err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	s := ctx.DeviceList.Find(uint32(deviceID))
	if s == nil {
		http.Error(w, fmt.Sprintf("No device with ID %d", deviceID), http.StatusNotFound)
		return
	}

	granted, err := s.Swipe(request.CardNumber, request.Door)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to emulate 'door open' (%v)", err), http.StatusInternalServerError)
		return
	}

	message := "Access denied"
	if granted {
		message = "Access granted"
	}

	response := struct {
		Granted bool   `json:"access-granted"`
		Message string `json:"message"`
	}{
		Granted: granted,
		Message: message,
	}

	b, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error generating response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func door(ctx *simulator.Context, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, fmt.Sprintf("Invalid method:%s - expected POST", r.Method), http.StatusMethodNotAllowed)
		return
	}

	url := r.URL.Path
	matches := regexp.MustCompile("^/uhppote/simulator/([0-9]+)/door/([1-4])$").FindStringSubmatch(url)
	deviceID, err := strconv.ParseUint(matches[1], 10, 32)
	if err != nil {
		http.Error(w, "Error reading request", http.StatusInternalServerError)
		return
	}

	door, err := strconv.ParseUint(matches[2], 10, 8)
	if err != nil {
		http.Error(w, "Error reading request", http.StatusInternalServerError)
		return
	}

	blob, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request", http.StatusInternalServerError)
		return
	}

	request := struct {
		Action string `json:"action"`
	}{}

	err = json.Unmarshal(blob, &request)
	if err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	switch request.Action {
	case "open":
		open(ctx, w, uint32(deviceID), uint8(door), blob)

	case "close":
		close(ctx, w, uint32(deviceID), uint8(door), blob)

	case "button":
		button(ctx, w, uint32(deviceID), uint8(door), blob)

	default:
		http.Error(w, fmt.Sprintf("%d: door %d, invalid action '%s'", deviceID, door, request.Action), http.StatusBadRequest)
	}
}

func open(ctx *simulator.Context, w http.ResponseWriter, deviceID uint32, door uint8, blob []byte) {
	request := struct {
		Duration uint `json:"duration"`
	}{}

	if err := json.Unmarshal(blob, &request); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	var duration *time.Duration
	if request.Duration > 0 {
		d := time.Duration(request.Duration) * time.Second
		duration = &d
	}

	s := ctx.DeviceList.Find(uint32(deviceID))
	if s == nil {
		http.Error(w, fmt.Sprintf("No device with ID %d", deviceID), http.StatusNotFound)
		return
	}

	opened, err := s.Open(door, duration)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to emulate 'door open' (%v)", err), http.StatusInternalServerError)
		return
	}

	message := "Could not open door"
	if opened {
		message = "Door open"
	}

	response := struct {
		Result  bool   `json:"result"`
		Message string `json:"message"`
	}{
		Result:  opened,
		Message: message,
	}

	b, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error generating response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func close(ctx *simulator.Context, w http.ResponseWriter, deviceID uint32, door uint8, blob []byte) {
	s := ctx.DeviceList.Find(uint32(deviceID))
	if s == nil {
		http.Error(w, fmt.Sprintf("No device with ID %d", deviceID), http.StatusNotFound)
		return
	}

	closed, err := s.Close(door)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to emulate 'door closed' (%v)", err), http.StatusInternalServerError)
		return
	}

	message := "Could not close door"
	if closed {
		message = "Door closed"
	}

	response := struct {
		Result  bool   `json:"result"`
		Message string `json:"message"`
	}{
		Result:  closed,
		Message: message,
	}

	b, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error generating response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func button(ctx *simulator.Context, w http.ResponseWriter, deviceID uint32, door uint8, blob []byte) {
	request := struct {
		Duration uint `json:"duration"`
	}{}

	if err := json.Unmarshal(blob, &request); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	duration := time.Duration(request.Duration) * time.Second

	s := ctx.DeviceList.Find(uint32(deviceID))
	if s == nil {
		http.Error(w, fmt.Sprintf("No device with ID %d", deviceID), http.StatusNotFound)
		return
	}

	unlocked, err := s.ButtonPressed(uint8(door), duration)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to emulate 'door button' (%v)", err), http.StatusInternalServerError)
		return
	}

	message := "Could not unlock door"
	if unlocked {
		message = "Door unlocked"
	}

	response := struct {
		Result  bool   `json:"result"`
		Message string `json:"message"`
	}{
		Result:  unlocked,
		Message: message,
	}

	b, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error generating response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
