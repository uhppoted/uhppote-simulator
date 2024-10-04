# CHANGELOG

## Unreleased

### Added
1. Added support for event auto-send and updated _get-event-listener_ and _set-event-listener_
   with the auto-send interval field.


## [0.8.9](https://github.com/uhppoted/uhppote-simulator/releases/tag/v0.8.9) - 2024-09-06

### Added
1. Implemented controller TCP/IP interface emulation.
2. REST _put-card_ API to add/update simulated controller cards.

### Updated
1. Reworked event handling to use connected UDP socket to send events.
2. Updated to Go 1.23.
3. Reworked `put-card` handler to return false if card _from_ or _to_ date is a zero value.


## [0.8.8](https://github.com/uhppoted/uhppote-simulator/releases/tag/v0.8.8) - 2024-03-27

### Added
1. `restore-default-parameters` command emulation.
2. Python REST API CLI client implementation.
3. Added public Docker image to ghcr.io.

### Updated
1. Bumped Go version to 1.22
2. Updated README with REST API.
3. Updated REST API OpenAPI specification.
4. Fixed uninitialised map in controller doors deserialization.


## [0.8.7](https://github.com/uhppoted/uhppote-simulator/releases/tag/v0.8.7) - 2023-12-01

### Added
1. `set-super-passwords` command and emulation.


## [0.8.6](https://github.com/uhppoted/uhppote-simulator/releases/tag/v0.8.6) - 2023-08-30

### Added
1. `activate-keypads` command and emulation.

### Updated
1. Replaced `nil` event pointer with zero value in `get-status`.


## [0.8.5](https://github.com/uhppoted/uhppote-simulator/releases/tag/v0.8.5) - 2023-06-13

### Added
1. `set-interlock` command and emulation.

### Updated
1. Replaced card `From` and `To` field pointers with zero values.


## [0.8.4](https://github.com/uhppoted/uhppote-simulator/releases/tag/v0.8.4) - 2023-03-17

### Added
1. `doc.go` package overview documentation.
2. `set-pc-control` command and emulation.
3. Added PIN to card record

### Updated
1. Replaced `math/rand` with `crypto/rand` for MAC address in create-device.


## [0.8.3](https://github.com/uhppoted/uhppote-simulator/releases/tag/v0.8.3) - 2022-12-16

### Added
1. Added ARM64 to release build artifacts

### Changed
1. Initialised `EventsList` in simulator default constructor (cf. https://github.com/uhppoted/uhppote-simulator/issues/6)
2. Reworked `EventsList` unmarshalling from JSON (cf. https://github.com/uhppoted/uhppote-simulator/issues/6)
   - Replaced zero values for `EventList` size and chunk with defaults 
   - Added check for zero chunk size before truncating
   - Reworked truncation to use calculated offset rather than loop
3. Reworked `checkTimeProfile` to include the controller time offset (cf. https://github.com/uhppoted/uhppote-simulator/issues/5)
4. Removed _zip_ files from release artifacts (no longer necessary)


## [0.8.2](https://github.com/uhppoted/uhppote-simulator/releases/tag/v0.8.2) - 2022-10-14

### Changed
1. Maintenance release for compatiblity with [uhppote-core](https://github.com/uhppoted/uhppote-core) v0.8.2

## [0.8.1](https://github.com/uhppoted/uhppote-simulator/releases/tag/v0.8.1) - 2022-01-01

### Changed
1. Maintenance release for compatiblity with [uhppote-core](https://github.com/uhppoted/uhppote-core) v0.8.1


## [0.8.0](https://github.com/uhppoted/uhppote-simulator/releases/tag/v0.8.0)

### Changed
1. Maintenance release for compatiblity with [uhppote-core](https://github.com/uhppoted/uhppote-core) v0.8.0


## [0.7.3](https://github.com/uhppoted/uhppote-simulator/releases/tag/v0.7.3)

### Changed
1. Maintenance release for compatiblity with [uhppote-core](https://github.com/uhppoted/uhppote-core) v0.7.3


## [0.7.2](https://github.com/uhppoted/uhppote-simulator/releases/tag/v0.7.2)

### Changed
1. Reworked the event list as a static array that discards a 'chunk' of events at the start
   of the array when the array is full. This matches the observed behaviour of a real-life
   UHPPOTE controller.
2. Updated `get-event` handler to return _overwritten_ if the requested event index is
   less than the _first_ event index in the stored event list.


## [0.7.1](https://github.com/uhppoted/uhppote-simulator/releases/tag/v0.7.1)

### Changed
1. Added handler for `clear-task-list`
2. Added handler for  `add-task`
3. Added handler for  `refresh-task-list`
4. Implemented task list emulation

## Older

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
