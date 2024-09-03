package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"

	"github.com/uhppoted/uhppote-core/types"

	"github.com/uhppoted/uhppote-simulator/simulator"
)

func cards(ctx *simulator.Context, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, fmt.Sprintf("Invalid method:%s - expected PUT", r.Method), http.StatusMethodNotAllowed)
		return
	}

	url := r.URL.Path
	matches := regexp.MustCompile("^/uhppote/simulator/([0-9]+)/cards/([0-9]+)$").FindStringSubmatch(url)

	controller, err := strconv.ParseUint(matches[1], 10, 32)
	if err != nil {
		http.Error(w, "Error reading request", http.StatusInternalServerError)
		return
	}

	card, err := strconv.ParseUint(matches[2], 10, 32)
	if err != nil {
		http.Error(w, "Error reading request", http.StatusInternalServerError)
		return
	}

	blob, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request", http.StatusInternalServerError)
		return
	}

	request := struct {
		Start string  `json:"start-date"`
		End   string  `json:"end-date"`
		Doors []uint8 `json:"doors"`
		PIN   uint32  `json:"PIN"`
	}{}

	err = json.Unmarshal(blob, &request)
	if err != nil {
		http.Error(w, "Invalid PUT card request", http.StatusBadRequest)
		return
	}

	from, _ := types.ParseDate(request.Start)
	to, _ := types.ParseDate(request.End)

	s := ctx.DeviceList.Find(uint32(controller))
	if s == nil {
		http.Error(w, fmt.Sprintf("No controller with ID %d", controller), http.StatusNotFound)
		return
	}

	if err := s.StoreCard(uint32(card), from, to, request.Doors, request.PIN); err != nil {
		http.Error(w, fmt.Sprintf("Error storing card %v (%v)", card, err), http.StatusInternalServerError)
		return
	}

	http.Error(w, fmt.Sprintf("card %v added/updated", card), http.StatusOK)
}
