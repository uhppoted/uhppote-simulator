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
go build -o bin ./...
```

#### Dependencies

| *Dependency*                                             | *Description*                                          |
| -------------------------------------------------------- | ------------------------------------------------------ |
| [uhppote-core](https://github.com/uhppoted/uhppote-core) | Device level API implementation                        |
| golang.org/x/lint/golint                                 | Additional *lint* check for release builds             |

## uhppote-simulator

Supported functions:
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
- GetEvent
- GetEventIndex
- SetEventIndex
- OpenDoor
- Listen

Usage: *uhppote-simulator \<command\> --devices=\<dir\>*

Defaults to 'run' unless one of the commands below is specified: 

- help
- version

Supported options:
- --bind <IP address to bind to>
- --devices <directory path for device files>
- --debug



