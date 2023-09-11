package UT0311L04

import (
    "net"
    "reflect"
    "testing"

    "github.com/uhppoted/uhppote-core/messages"
    "github.com/uhppoted/uhppote-simulator/entities"
)

func TestSetSuperPasswords(t *testing.T) {
    txq := make(chan entities.Message, 8)

    s := UT0311L04{
        SerialNumber: 405419896,
        Doors: entities.MakeDoors(),

        txq: txq,
    }

    src := net.UDPAddr{IP: net.IPv4(10, 0, 0, 1), Port: 12345}

    expected := struct {
        response entities.Message
    }{
      response:  entities.Message{
        Destination: &src,
        Message: &messages.SetSuperPasswordsResponse{
            SerialNumber: 405419896,
            Succeeded:    true,
        },
    },
}

    request := messages.SetSuperPasswordsRequest{
        SerialNumber: 405419896,
        Door:    3,
        Password1: 12345,
        Password2: 0,
        Password3: 999999,
        Password4: 54321,
    }

    s.SetSuperPasswords(&src, &request)

    response := <-txq

    if !reflect.DeepEqual(response, expected.response) {
        t.Errorf("'set-super-passwords' sent incorrect response\n   expected: %+v\n   got:      %+v\n", expected.response, response)
    }

    // if !reflect.DeepEqual(s.Doors, expected.doors) {
    //     t.Errorf("'set-super-passwords' failed to update simulator\n   expected: %+v\n   got:      %+v\n", true, s.Doors)
    // }
}

