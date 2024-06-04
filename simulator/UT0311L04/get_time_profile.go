package UT0311L04

import (
	"time"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-core/types"
)

func (s *UT0311L04) getTimeProfile(request *messages.GetTimeProfileRequest) (*messages.GetTimeProfileResponse, error) {
	if s.SerialNumber != request.SerialNumber {
		return nil, nil
	}

	response := messages.GetTimeProfileResponse{
		SerialNumber: s.SerialNumber,
	}

	if request.ProfileID > 1 && request.ProfileID < 255 {
		if profile, ok := s.TimeProfiles[request.ProfileID]; ok {
			// FIXME: replace From in types.TimeProfile
			from := types.Date{}
			if profile.From != nil {
				from = *profile.From
			}

			// FIXME: replace To in TimeProfile entity
			to := types.Date{}
			if profile.To != nil {
				to = *profile.To
			}

			response.ProfileID = profile.ID
			response.LinkedProfileID = profile.LinkedProfileID
			response.From = from
			response.To = to

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

	return &response, nil
}
