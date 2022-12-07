# CHANGELOG

## [Unreleased]

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
