![build](https://github.com/uhppoted/uhppote-simulator/workflows/build/badge.svg)
![ghcr](https://github.com/uhppoted/uhppote-simulator/workflows/ghcr/badge.svg)

# uhppote-simulator

Software simulator for access control systems based on the *UHPPOTE UT0311-L0x* TCP/IP Wiegand access control boards. 

Supported operating systems:
- Linux
- MacOS
- Windows
- RaspberryPi (ARM/ARM7/ARM6)

---
### Contents

- [Release Notes](#release-notes)
- [Installation](#installation)
   - [Docker](#docker)
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

**[v0.8.11](https://github.com/uhppoted/uhppote-simulator/releases/tag/v0.8.11) - 2025-07-01**

1. Added `get/set-antipassback` command emulation.
2. Fixed incoming requests not printed to console (bug).
3. Updated to Go 1.24.


## Installation

### Docker

A public _Docker_ image is published to [ghcr.io](https://github.com/uhppoted?tab=packages&repo_name=uhppote-simulator). The image:
- includes a _getting started_ sample controller\
- is configured to store emulated controllers in /usr/local/etc/uhppoted/simulator

#### `docker compose`

A sample Docker `compose` configuration is provided in the [`docker/compose`](docker/compose) folder. 

To run the example, download and extract the [zipped](docker/compose.zip) scripts and supporting files into folder
of your choice and then:
```
cd <compose folder>
docker compose up
```

The emulated controllers can be managed using the [REST API](#rest-api).

#### `docker run`

To start a simulator using Docker `run`:
```
docker pull ghcr.io/uhppoted/simulator:latest
docker run --detach --publish 8000:8000 --publish 60000:60000/udp --name simulator \
           --mount source=uhppoted,target=/usr/local/etc/uhppoted \
           --rm ghcr.io/uhppoted/simulator
```

The emulated controllers can be managed using the REST API [REST API](#rest-api).

#### `docker build`

For inclusion in a Dockerfile:
```
FROM ghcr.io/uhppoted/simulator:latest
```

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
- GetAntiPassback
- SetAntiPassback
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
- using the [REST.py](scripts) script
- using one of the many [Postman-like](https://www.postman.com/) tools available. Postman scripts can be found in
  the [scripts](scripts) folder
- using the [Swagger Editor](https://editor.swagger.io) with the [OpenAPI](https://github.com/uhppoted/uhppote-simulator/blob/main/documentation/simulator-api.yaml) YAML file.

The default port is 8000.

### `swipe-card`
Simulates a card swipe with optional keypad code.
```
URL: `http://localhost:8000/uhppote/simulator/{controller}/swipe`
Method: POST
Request:
{
    "door": <door>,
    "card-number: <card>,
    "direction": [1|2],
    "PIN": <passcode>
}

controller   controller serial number e.g. 405419896
door         door number [1..4] e.g. 3
direction    1: IN, 2: OUT
PIN          (optional) PIN code for keypad reader
```
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
```
URL: `http://localhost:8000/uhppote/simulator/{controller}/code`
Method: POST
Request:
{
    "door": <door>,
    "PIN": <passcode>
}

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
```
URL: `http://localhost:8000/uhppote/simulator/{controller}/door/{door}`
Method: POST
Request:
{
    "action": "open",
}

controller   controller serial number e.g. 405419896
door         door number [1..4] e.g. 3
action       open
```
```
curl -X POST "http://127.0.0.1:8000/uhppote/simulator/405419896/door/1" -H "accept: application/json" -H "Content-Type: application/json" -d '{"action":"open"}'
```

### `close-door`
Simulates closing a door - a door closed event will be generated if the door was open.
```
URL: `http://localhost:8000/uhppote/simulator/{controller}/door/{door}`
Method: POST
Request:
{
    "action": "close",
}

controller   controller serial number e.g. 405419896
door         door number [1..4] e.g. 3
action       close
```
```
curl -X POST "http://127.0.0.1:8000/uhppote/simulator/405419896/door/1" -H "accept: application/json" -H "Content-Type: application/json" -d '{"action":"close"}'
```

### `press-button`
Simulates pressing the 'door open' button - a `button pressed` event will be generated if the button was not already pressed.
```
URL: `http://localhost:8000/uhppote/simulator/{controller}/door/{door}`
Method: POST
Request:
{
    "action": "button",
    "duration": 10
}

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
```
URL: `http://localhost:8000/uhppote/simulator`
Method: GET
```
```
curl -X GET "http://127.0.0.1:8000/uhppote/simulator" -H "accept: application/json"
```

### `create-controller`
Adds a new controller to the simulator.
```
URL: `http://localhost:8000/uhppote/simulator`
Method: POST
Request:
{
    "device-id": <controller>,
    "device-type": <manufacturer-code>,
    "compressed": [true|false]
}

controller          controller serial number e.g. 405419896
manufacturer-code   UT0311-L02 for a 2 door controller, UT0311-L04 for a 4 door controller
compressed          store controller in compressed (true) or human-readable form (false)
```
```
curl -X POST "http://127.0.0.1:8000/uhppote/simulator" -H "accept: */*" -H "Content-Type: application/json" -d '{"device-id":405419896,"device-type":"UT0311-L04","compressed":false}'
```

### `delete-controller`
Delets a controller from the simulator.
```
URL: `http://localhost:8000/uhppote/simulator/<controller>`
Method: DELETE

controller   controller serial number e.g. 405419896
```
```
curl -X DELETE "http://127.0.0.1:8000/uhppote/simulator/405419896" -H "accept: */*"
```

### `put-card`
Adds or updates a card on a simulated controller. Unlike the controller `put-card` API, the REST API does not 
require a valid start/end date (for testing purposes).
```
URL: `http://localhost:8000/uhppote/simulator/<controller>/cards/<card>`
Method: PUT
Request:
{
    "start-date": <start-date>,
    "end-date": <end-date>,
    "doors": <doors>,
    "PIN": <uint32>
}

controller     controller serial number e.g. 405419896
start-date     card 'valid from' date e.g 2024-01-01
end-date       card 'valid until' date e.g 2024-12-31
doors          list of doors the card for which the card has access e.g. [1,2,4]
PIN            card PIN e.g.7531
```
```
curl -X POST "http://127.0.0.1:8000/uhppote/simulator/405419896/cards/10058400" -H "accept: */*" -H "Content-Type: application/json" -d '{"start-date": "2024-01-01", "end-date": "2024-12-31", "doors": [1,2,4], "PIN": 7531}'
```


### Postman

### Python



## NOTES

### `put-card`

1. The UHPPOTE access controller has a weird behaviour around the PIN field. According to the SDK 
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

2. The emulated controller does not accept cards with a 'zero' start or end date - the REST _put-card_ API 
can be used to add cards without valid start/end dates for testing.


### `restore-default-parameters`

`restore-default-parameters` has (for practical reasons) NOT been validated against an actual controller. Resetting
the simulator:
- clears the internal controller IPv4 address, netmask and gateway
- clears the event listener address
- clears all events
- deletes all cards
- sets all doors to `controlled` mode and 5 seconds delay
- clears all door passcodes
- sets the anti-passback mode to _disabled_


### `passcode`

1. If a supervisor passcode is entered for a door that is _normally closed_, the UHPPOTE controller will
   unlock the door and then immediately relock it. This seems anomalous and in this case the simulator 
   unlocks the door on the assumption that the supervisor code is intended to be an override.


### Events

_tl;dr; The UHPPOTE controller does not 'rollover' when the onboard event store is filled._

From experimentation, it appears that the UHPPOTE controller has an event store for approximately 200 000 events
(the user manual says 100 000, so possibly varies with model/version). On filling the event buffer the controller
seems to discard a _chunk_ of about 2048 events from the start of the event buffer to make space for new events. 
The event index continues to increment monotonically (presumably until the uint32 overflows and wraps back to 0).



