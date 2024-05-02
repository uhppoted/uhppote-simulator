package UT0311L04

import (
	"fmt"

	"github.com/uhppoted/uhppote-core/messages"
)

func (s *UT0311L04) clearTimeProfiles(request *messages.ClearTimeProfilesRequest) (*messages.ClearTimeProfilesResponse, error) {
	if s.SerialNumber != request.SerialNumber {
		return nil, nil
	}

	cleared := false

	if request.MagicWord == 0x55aaaa55 {
		cleared = s.TimeProfiles.Clear()
	}

	response := messages.ClearTimeProfilesResponse{
		SerialNumber: s.SerialNumber,
		Succeeded:    cleared,
	}

	if err := s.Save(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	return &response, nil
}
