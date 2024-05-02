package UT0311L04

import (
	"fmt"
	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-core/types"
	"github.com/uhppoted/uhppote-simulator/entities"
	"net"
	"time"
)

func (s *UT0311L04) setTime(addr *net.UDPAddr, request *messages.SetTimeRequest) (*messages.SetTimeResponse, error) {
	if s.SerialNumber != request.SerialNumber {
		return nil, nil
	}

	dt := time.Time(request.DateTime).Format("2006-01-02 15:04:05")
	if utc, err := time.ParseInLocation("2006-01-02 15:04:05", dt, time.UTC); err != nil {
		return nil, err
	} else {
		now := time.Now().UTC()
		delta := utc.Sub(now)
		datetime := now.Add(delta)

		s.TimeOffset = entities.Offset(delta)
		response := messages.SetTimeResponse{
			SerialNumber: s.SerialNumber,
			DateTime:     types.DateTime(datetime),
		}

		if err = s.Save(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}

		return &response, nil
	}
}
