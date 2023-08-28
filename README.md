 ![build](https://github.com/uhppoted/uhppote-simulator/workflows/build/badge.svg)

# uhppote-simulator

Software simulator for access control systems based on the *UHPPOTE UT0311-L0x* TCP/IP Wiegand access control boards. 

Supported operating systems:
- Linux
- MacOS
- Windows
- ARM7 (RaspberryPi)

## Releases

| *Version* | *Description*                                                                             |
| --------- | ----------------------------------------------------------------------------------------- |
| v0.8.6    | Implemented reader keypas emulation and `activate-keypads` API                              |
| v0.8.5    | Implemented door interlock emulation and `set-interlock` API                              |
| v0.8.4    | Added `set-pc-control` emulation and support for card keypad PINs                         |
| v0.8.3    | Bug for time profiles timezone handling and events list JSON marshalling                  |
| v0.8.2    | Maintenance release for version compatibility with `uhppote-core` v0.8.2                  |
| v0.8.1    | Maintenance release for version compatibility with `uhppote-core` v0.8.1                  |
| v0.8.0    | Maintenance release for version compatibility with `uhppote-core` v0.8.0                  |
| v0.7.3    | Maintenance release for version compatibility with `uhppote-core` v0.7.3                  |
| v0.7.2    | Replaced event list rollover to discard a 'chunk' of events when the list is full         |
| v0.7.1    | Added support for task list functionality from the extended API                           |
| v0.7.0    | Added support for time profiles from the extended API                                     |
| v0.6.12   | Added simulation support for encoding `nil` events for response to `get-status`           |
| v0.6.10   | Maintenance release for version compatibility with `uhppoted-app-wild-apricot`            |
| v0.6.8    | Implements firmware _UHPPOTE_ v6.62 event message format                                  |
| v0.6.7    | Implements `record-special-events` and door open, close and button actions and events     |
| v0.6.5    | Maintenance release for version compatibility with `node-red-contrib-uhppoted`            |
| v0.6.4    | Maintenance release for version compatibility with `uhppoted-app-sheets`                  |
| v0.6.3    | Reworked card list as fixed length array and emulated deleted records                     |
| v0.6.2    | Updated simulation response for 'no events'                                               |
| v0.6.1    | Maintenance release to update module dependencies                                         |
| v0.6.0    | Maintenance release to update module dependencies                                         |
| v0.5.1    | Initial release following restructuring into standalone Go *modules* and *git submodules* |

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

### Events

_tl;dr; The UHPPOTE controller does not 'rollover' when the onboard event store is filled._

From experimentation, it appears that the UHPPOTE controller has an event store for approximately 200 000 events
(possibly varies varies with model/version). On filling the event buffer the controller seems to discard a _chunk_
of about 2048 events from the start of the event buffer to make space for new events. The event index continues to
increment monotonically.



