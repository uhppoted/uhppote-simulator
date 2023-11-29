 ![build](https://github.com/uhppoted/uhppote-simulator/workflows/build/badge.svg)

# uhppote-simulator

Software simulator for access control systems based on the *UHPPOTE UT0311-L0x* TCP/IP Wiegand access control boards. 

Supported operating systems:
- Linux
- MacOS
- Windows
- ARM7 (RaspberryPi)

## Release Notes

### Current release

**[v0.8.7](https://github.com/uhppoted/uhppote-simulator/releases/tag/v0.8.7) - 2023-12-01**

1. ``set-super-passwords` command and emulation.

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

#### Dependencies

| *Dependency*                                             | *Description*                                          |
| -------------------------------------------------------- | ------------------------------------------------------ |
| [uhppote-core](https://github.com/uhppoted/uhppote-core) | Device level API implementation                        |


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
- Listen

Supported _actions_:
- swipe card
- press button
- open door

Usage: *uhppote-simulator \<command\> --devices=\<dir\>*

Defaults to 'run' unless one of the commands below is specified: 

- help
- version

Supported options:
- --bind <IP address to bind to>
- --devices <directory path for device files>
- --debug

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



