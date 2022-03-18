## CHANGELOG

### v0.7.2

1. Reworked the event list as a static array that discards a 'chunk' of events at the start
   of the array when the array is full. This matches the observed behaviour of a real-life
   UHPPOTE controller.
2. Updated `get-event` handler to return _overwritten_ if the requested event index is
   less than the _first_ event index in the stored event list.

### v0.7.1

1. Added handler for `clear-task-list`
2. Added handler for  `add-task`
3. Added handler for  `refresh-task-list`
4. Implemented task list emulation