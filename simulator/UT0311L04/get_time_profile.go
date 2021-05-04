package UT0311L04

import (
	"net"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-core/types"
)

func (s *UT0311L04) getTimeProfile(addr *net.UDPAddr, request *messages.GetTimeProfileRequest) {
	if s.SerialNumber == request.SerialNumber {
		response := messages.GetTimeProfileResponse{
			SerialNumber: s.SerialNumber,
		}

		if request.ProfileID > 1 && request.ProfileID < 255 {
			if profile, ok := s.TimeProfiles[request.ProfileID]; ok {

				response.ProfileID = profile.ID
				response.LinkedProfileID = profile.LinkedProfileID
				response.From = profile.From
				response.To = profile.To

				response.Monday = profile.Weekdays[types.Monday]
				response.Tuesday = profile.Weekdays[types.Tuesday]
				response.Wednesday = profile.Weekdays[types.Wednesday]
				response.Thursday = profile.Weekdays[types.Thursday]
				response.Friday = profile.Weekdays[types.Friday]
				response.Saturday = profile.Weekdays[types.Saturday]
				response.Sunday = profile.Weekdays[types.Sunday]

				if segment, ok := profile.Segments[1]; ok {
					response.Segment1Start = segment.Start
					response.Segment1End = segment.End
				}

				if segment, ok := profile.Segments[2]; ok {
					response.Segment2Start = segment.Start
					response.Segment2End = segment.End
				}

				if segment, ok := profile.Segments[3]; ok {
					response.Segment3Start = segment.Start
					response.Segment3End = segment.End
				}
			}
		}

		s.send(addr, &response)
	}
}
