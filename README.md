 ![build](https://github.com/uhppoted/uhppote-simulator/workflows/build/badge.svg)

# uhppote-simulator

Software simulator for access control systems based on the *UHPPOTE UT0311-L0x* TCP/IP Wiegand access control boards. 

Supported operating systems:
- Linux
- MacOS
- Windows
- ARM7 (RaspberryPi)

---
### Contents

- [Release Notes](#release-notes)
- [Installation](#installation)
   - [Building from source](#building-from-source)
- [Supported functions](#uhppote-simulator)
- [REST API](#rest-api)
   - [_swipe-card_](#swipe-card)
   - [_supervisor-passcode_](#supervisor-passcode)
   - [_open-door_](#open-door)
   - [_close-door_](#close-door)
   - [_press-button_](#press-button)
   - [_list-controllers_](#list-controllers)
   - [_create-controller_](#create-controller)
   - [_delete-controller_](#delete-controller)
- [Notes](#notes)
   - [`put-card`](#put-card)
   - [`restore-default-parameters`](#restore-default-parameters)
   - [`passcode`](#passcode)

---

## Release Notes

### Current release

**[v0.8.7](https://github.com/uhppoted/uhppote-simulator/releases/tag/v0.8.7) - 2023-12-01**

1. `set-super-passwords` command and emulation.

## Installation

### Building from source

Assuming you have `Go` and `make` installed:

```
git clone https://github.com/uhppoted/uhppote-simulator.git
cd uhppote-simulator
make build
```

If you prefer not to use `make`:
```
git clone https://github.com/uhppoted/uhppote-simulator.git
cd uhppote-simulator
mkdir bin
go build -trimpath -o bin ./...
```

## uhppote-simulator

Supported `uhppote` functions:
- FindDevices
- FindDevice
- SetAddress
- GetListener
- SetListener
- GetTime
- SetTime
- GetDoorControlState
- SetDoorControlState
- RecordSpecialEvents
- GetStatus
- GetCards
- GetCardById
- GetCardByIndex
- PutCard
- DeleteCard
- DeleteCards
- GetTimeProfile
- SetTimeProfile
- ClearTimeProfiles
- ClearTaskList
- AddTask
- RefreshTaskList
- GetEvent
- GetEventIndex
- SetEventIndex
- OpenDoor
- SetPCControl
- SetInterlock
- ActivateAccessKeypads
- RestoreDefaultParameters
- Listen

Usage: *uhppote-simulator \<command\> --devices=\<dir\>*

Defaults to 'run' unless one of the commands below is specified: 

- help
- version

Supported options:
- --bind <IP address to bind to>
- --devices <directory path for device files>
- --debug

## REST API

The simulator provides a REST API to simulate user actions and manage controllers:

- swipe card
- supervisor passcode
- open door
- close door
- press button
- list controllers
- create controller
- delete controller

The actions may be invoked:
- from the command line using _curl_ 
- using one of the many [Postman-like](https://www.postman.com/) tools available
- using the [Swagger Editor](https://editor.swagger.io) with the [OpenAPI](https://github.com/uhppoted/uhppote-simulator/blob/main/documentation/simulator-api.yaml) YAML file.

The default port is 8000.

### `swipe-card`

Simulates a card swipe with optional keypad code.

URL: `http://localhost:8000/uhppote/simulator/{controller}/swipe`
Method: POST
Request:
```
{
    "door": <door>,
    "card-number: <card>,
    "direction": [1|2],
    "PIN": <passcode>
}
```

`controller`   controller serial number e.g. 405419896
`door`         door number [1..4] e.g. 3
`direction`    1: IN, 2: OUT
`PIN`          (optional) PIN code for keypad reader

swipe:
```
curl -X POST "http://127.0.0.1:8000/uhppote/simulator/405419896/swipe" -H "accept: application/json" -H "Content-Type: application/json" -d '{"door":1,"card-number":10058400,"direction":1,"PIN":1357}'
```
swipe-in:
```
curl -X POST "http://127.0.0.1:8000/uhppote/simulator/405419896/swipe" -H "accept: application/json" -H "Content-Type: application/json" -d '{"door":1,"card-number":10058400,"direction":1,"PIN":1357}'
```
swipe-out:
```
curl -X POST "http://127.0.0.1:8000/uhppote/simulator/405419896/swipe" -H "accept: application/json" -H "Content-Type: application/json" -d '{"door":1,"card-number":10058400,"direction":2,"PIN":1357}'
```

### `supervisor-passcode`

Simulates a supervisor passcode entry on a keypad.

URL: `http://localhost:8000/uhppote/simulator/{controller}/code`
Method: POST
Request:
```
{
    "door": <door>,
    "PIN": <passcode>
}
```
```
controller   controller serial number e.g. 405419896
door         door number [1..4] e.g. 3
passcode     supervisor passcode
```
```
curl -X POST "http://127.0.0.1:8000/uhppote/simulator/405419896/code" -H "accept: application/json" -H "Content-Type: application/json" -d '{"door":1,"passcode":1357}'
```

### `open-door`

Simulates opening a door - the door must be unlocked (e.g. by a card swipe or button press or supervisor passcode) to open. A `door open` event will generated if the
door was closed.

URL: `http://localhost:8000/uhppote/simulator/{controller}/door/{door}`
Method: POST
Request:
```
{
    "action": "open",
}
```
```
controller   controller serial number e.g. 405419896
door         door number [1..4] e.g. 3
action       open
```
```
curl -X POST "http://127.0.0.1:8000/uhppote/simulator/405419896/door/1" -H "accept: application/json" -H "Content-Type: application/json" -d '{"action":"open"}'
```

### `close-door`

Simulates closing a door - a door closed event will be generated if the door was open.

URL: `http://localhost:8000/uhppote/simulator/{controller}/door/{door}`
Method: POST
Request:
```
{
    "action": "close",
}
```
```
controller   controller serial number e.g. 405419896
door         door number [1..4] e.g. 3
action       close
```
```
curl -X POST "http://127.0.0.1:8000/uhppote/simulator/405419896/door/1" -H "accept: application/json" -H "Content-Type: application/json" -d '{"action":"close"}'
```

### `press-button`

Simulates pressing the 'door open' button - a `button pressed` event will be generated if the button was not already pressed.

URL: `http://localhost:8000/uhppote/simulator/{controller}/door/{door}`
Method: POST
Request:
```
{
    "action": "button",
    "duration": 10
}
```
```
controller   controller serial number e.g. 405419896
door         door number [1..4] e.g. 3
action       button
duration     seconds to press the button e.g. 10
```
```
curl -X POST "http://127.0.0.1:8000/uhppote/simulator/405419896/door/1" -H "accept: application/json" -H "Content-Type: application/json" -d '{"action":"button", "duration":10}'
```

### `list-controllers`

Lists the configured controllers.

URL: `http://localhost:8000/uhppote/simulator`
Method: GET

```
curl -X GET "http://127.0.0.1:8000/uhppote/simulator" -H "accept: application/json"
```

### `create-controller`

Adds a new controller to the simulator.

URL: `http://localhost:8000/uhppote/simulator`
Method: POST
Request:
```
{
    "device-id": <controller>,
    "device-type": <manufacturer-code>,
    "compressed": [true|false]
}
```
```
controller          controller serial number e.g. 405419896
manufacturer-code   UT0311-L02 for a 2 door controller, UT0311-L04 for a 4 door controller
compressed          store controller in compressed (true) or human-readable form (false)
```
```
curl -X POST "http://127.0.0.1:8000/uhppote/simulator" -H "accept: */*" -H "Content-Type: application/json" -d '{"device-id":405419896,"device-type":"UT0311-L04","compressed":false}'
```

### `delete-controller`

Delets a controller from the simulator.

URL: `http://localhost:8000/uhppote/simulator/<controller>`
Method: DELETE
```
controller   controller serial number e.g. 405419896
```
```
curl -X DELETE "http://127.0.0.1:8000/uhppote/simulator/405419896" -H "accept: */*"
```

### Postman

### Python



## NOTES

### `put-card`

The UHPPOTE access controller has a weird behaviour around the PIN field. According to the SDK 
documentation, valid PINs are in the range 0 to 999999. However the controller will accept a 
PIN number out of that range and only keep the lower 7 nibbles of the 32-bit unsigned value.
e.g:

| PIN     | Hex value | Stored as (hex) | Retrieved as (hex) | Retrieved as (decimal) |
|---------|-----------|-----------------|--------------------|------------------------|
| 0       | 0x000000  | 0x000000        | 0x000000           | 0                      |
| 999999  | 0x0f423f  | 0x0f423f        | 0x0f423f           | 999999                 |
| 1000000 | 0x0f4240  | 0x000000        | 0x000000           | 0                      |
| 1000001 | 0x0f4241  | 0x000000        | 0x000000           | 0                      |
| 1048576 | 0x100000  | 0x000000        | 0x000000           | 0                      |
| 1048577 | 0x100001  | 0x000000        | 0x000001           | 1                      |
| 1999999 | 0x1E847F  | 0x0E847F        | 0x000001           | 951423                 |

Note that by design, the simulator does not emulate this behaviour, on the grounds that it is probably a 
version specific bug.

### `restore-default-parameters`

`restore-default-parameters` has (for practical reasons) NOT been validated against an actual controller. Resetting
the simulator:
- clears the internal controller IPv4 address, netmask and gateway
- clears the event listener address
- clears all events
- deletes all cards
- sets all doors to `controlled` mode and 5 seconds delay
- clears all door passcodes

### `passcode`

1. If a supervisor passcode is entered for a door that is _normally closed_, the UHPPOTE controller will
   unlock the door and then immediately relock it. This seems anomalous and in this case the simulator 
   unlocks the door on the assumption that the supervisor code is intended to be an override.

### Events

_tl;dr; The UHPPOTE controller does not 'rollover' when the onboard event store is filled._

From experimentation, it appears that the UHPPOTE controller has an event store for approximately 200 000 events
(possibly varies varies with model/version). On filling the event buffer the controller seems to discard a _chunk_
of about 2048 events from the start of the event buffer to make space for new events. The event index continues to
increment monotonically.



