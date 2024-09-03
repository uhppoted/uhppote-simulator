package UT0311L04

import (
	"reflect"
	"testing"

	"github.com/uhppoted/uhppote-core/messages"
	"github.com/uhppoted/uhppote-core/types"

	"github.com/uhppoted/uhppote-simulator/entities"
)

func TestHandlePutCardRequest(t *testing.T) {
	request := messages.PutCardRequest{
		SerialNumber: 12345,
		CardNumber:   192837465,
		From:         types.MustParseDate("2019-01-01"),
		To:           types.MustParseDate("2019-12-31"),
		Door1:        1,
		Door2:        0,
		Door3:        1,
		Door4:        0,
	}

	response := messages.PutCardResponse{
		SerialNumber: 12345,
		Succeeded:    true,
	}

	testHandle(&request, &response, t)
}

func TestHandlePutCardRequestWithZeroFromDate(t *testing.T) {
	request := messages.PutCardRequest{
		SerialNumber: 405419896,
		CardNumber:   10058400,
		From:         types.Date{},
		To:           types.MustParseDate("2024-12-31"),
		Door1:        1,
		Door2:        0,
		Door3:        1,
		Door4:        0,
	}

	expected := messages.PutCardResponse{
		SerialNumber: 405419896,
		Succeeded:    false,
	}

	ut0311L04 := UT0311L04{
		SerialNumber: 405419896,
		Cards:        entities.CardList{},
	}

	if response, err := ut0311L04.putCard(&request); err != nil {
		t.Fatalf("%v", err)
	} else if response == nil {
		t.Errorf("Invalid 'put-card' response: Expected: %v, got: %v", expected, response)
	} else if !reflect.DeepEqual(*response, expected) {
		t.Errorf("Incorrect 'put-card' response: Expected:\n%v, got:\n%v", expected, *response)
	}
}

func TestHandlePutCardRequestWithZeroToDate(t *testing.T) {
	request := messages.PutCardRequest{
		SerialNumber: 405419896,
		CardNumber:   10058400,
		From:         types.MustParseDate("2024-01-01"),
		To:           types.Date{},
		Door1:        1,
		Door2:        0,
		Door3:        1,
		Door4:        0,
	}

	expected := messages.PutCardResponse{
		SerialNumber: 405419896,
		Succeeded:    false,
	}

	ut0311L04 := UT0311L04{
		SerialNumber: 405419896,
		Cards:        entities.CardList{},
	}

	if response, err := ut0311L04.putCard(&request); err != nil {
		t.Fatalf("%v", err)
	} else if response == nil {
		t.Errorf("Invalid 'put-card' response: Expected: %v, got: %v", expected, response)
	} else if !reflect.DeepEqual(*response, expected) {
		t.Errorf("Incorrect 'put-card' response: Expected:\n%v, got:\n%v", expected, *response)
	}
}
