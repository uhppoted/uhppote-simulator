package entities

import (
	"github.com/uhppoted/uhppote-core/types"
)

type TimeProfiles map[uint8]types.TimeProfile

func (t TimeProfiles) Set(profile types.TimeProfile) bool {
	if profile.ID > 1 && profile.ID < 255 {
		t[profile.ID] = profile

		return true
	}

	return false
}

func (t TimeProfiles) Clear() bool {
	for k := range t {
		delete(t, k)
	}

	return true
}
