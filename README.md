![build](https://github.com/uhppoted/uhppote-simulator/workflows/build/badge.svg)

# uhppote-simulator

Software simulator for access control systems based on the *UHPPOTE UT0311-L0x* TCP/IP Wiegand access control boards. 

Supported operating systems:
- Linux
- MacOS
- Windows

## Releases

| *Version* | *Description*                                                                             |
| --------- | ----------------------------------------------------------------------------------------- |
| v0.6.3    | Reworked card list as fixed length array and emulated deleted records                     |
| v0.6.2    | Updated simulation response for 'no events'                                               |
| v0.6.1    | Maintenance release to update module dependencies                                         |
| v0.6.0    | Maintenance release to update module dependencies                                         |
| v0.5.1    | Initial release following restructuring into standalone Go *modules* and *git submodules* |

## Installation

### Building from source

#### Dependencies

| *Dependency*                                             | *Description*                                          |
| -------------------------------------------------------- | ------------------------------------------------------ |
| [uhppote-core](https://github.com/uhppoted/uhppote-core) | Device level API implementation                        |
| golang.org/x/lint/golint                                 | Additional *lint* check for release builds             |

### Binaries

## uhppote-simulator

Supported functions:
- FindDevices
- FindDevice
- SetAddress
- GetStatus
- GetTime
- SetTime
- GetDoorControlState
- SetDoorControlState
- GetListener
- SetListener
- GetCards
- GetCardByIndex
- GetCardById
- PutCard
- DeleteCard
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



