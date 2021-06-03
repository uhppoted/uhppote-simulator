package UT0311L04

import (
	"net"
	"time"

	"github.com/uhppoted/uhppote-core/messages"
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

				response.Monday = profile.Weekdays[time.Monday]
				response.Tuesday = profile.Weekdays[time.Tuesday]
				response.Wednesday = profile.Weekdays[time.Wednesday]
				response.Thursday = profile.Weekdays[time.Thursday]
				response.Friday = profile.Weekdays[time.Friday]
				response.Saturday = profile.Weekdays[time.Saturday]
				response.Sunday = profile.Weekdays[time.Sunday]

				if segment, ok := profile.Segments[1]; ok {
					response.Segment1Start = &segment.Start
					response.Segment1End = &segment.End
				}

				if segment, ok := profile.Segments[2]; ok {
					response.Segment2Start = &segment.Start
					response.Segment2End = &segment.End
				}

				if segment, ok := profile.Segments[3]; ok {
					response.Segment3Start = &segment.Start
					response.Segment3End = &segment.End
				}
			}
		}

		s.send(addr, &response)
	}
}
