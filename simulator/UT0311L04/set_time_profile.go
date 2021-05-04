package UT0311L04

import (
	"fmt"
	"net"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-core/types"
)

func (s *UT0311L04) setTimeProfile(addr *net.UDPAddr, request *messages.SetTimeProfileRequest) {
	if s.SerialNumber == request.SerialNumber {
		profile := types.TimeProfile{
			ID:              request.ProfileID,
			LinkedProfileID: request.LinkedProfileID,
			From:            &request.From,
			To:              &request.To,
			Weekdays: types.Weekdays{
				types.Monday:    request.Monday,
				types.Tuesday:   request.Tuesday,
				types.Wednesday: request.Wednesday,
				types.Thursday:  request.Thursday,
				types.Friday:    request.Friday,
				types.Saturday:  request.Saturday,
				types.Sunday:    request.Sunday,
			},
			Segments: types.Segments{
				1: types.Segment{
					Start: &request.Segment1Start,
					End:   &request.Segment1End,
				},
				2: types.Segment{
					Start: &request.Segment2Start,
					End:   &request.Segment2End,
				},
				3: types.Segment{
					Start: &request.Segment3Start,
					End:   &request.Segment3End,
				},
			},
		}

		ok := s.TimeProfiles.Set(profile)

		response := messages.SetTimeProfileResponse{
			SerialNumber: s.SerialNumber,
			Succeeded:    ok,
		}

		s.send(addr, &response)

		if err := s.Save(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}
