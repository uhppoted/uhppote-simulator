## v0.6.x

### IN PROGRESS

- [ ] Add 'open' REST API
- [ ] Add 'button' REST API
- [ ] Remove relay and input state from JSON
- [ ] Replace UTO311L04.TimeOffset with time zone
- [ ] Unit tests for EventList
- [ ] Check real device events list rollover

- [x] Implement record-special-events
- [x] Date (for get-device) should be the manufactured date i.e. fixed, not 'now'

## TODO

### simulator
- [ ] concurrent requests
- [ ] simulator-cli
- [ ] HTML
- [ ] httpd
- [ ] Rework simulator.run to use rx channels
- [ ] Reload simulator on device file change
- [ ] Implement JSON unmarshal to initialise default values
- [ ] Swagger UI
- [ ] Autodetect gzipped files (https://stackoverflow.com/questions/28309988/how-to-read-from-either-gzip-or-plain-text-reader-in-golang)

### Documentation

- [ ] godoc
- [ ] build documentation
- [ ] install documentation
- [ ] user manuals
- [ ] man/info page

### Other

1.  Integration tests
2.  Verify fields in listen events/status replies against SDK:
    - battery status can be (at least) 0x00, 0x01 and 0x04
