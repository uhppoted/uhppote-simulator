# Notes

1. Controller weirdness when handling of PINs greater than 999999
   - Between 1000000 and 1048576 PIN is set to 0
   - Seems to just drop the most signicant nibble when the PIN >= 1048577
   - 999999  => 0x0F423F (stored as 0x000000)
   - 1000000 => 0x0F4240 (stored as 0x000000)
   - 1000001 => 0x0F4241 (stored as 0x000000)
   - 1048576 => 0x100000 (stored as 0x000000)
   - 1048577 => 0x100001 (stored as 0x000001)
   - 1999999 => 0x1E847F (stored as 0x0E847F)
   - probably just short-circuited byte-by-byte (nibble-by-nibble ?) equals
   - or maybe just masks out the most significant nibble ?

2. https://blog.cloudflare.com/everything-you-ever-wanted-to-know-about-udp-sockets-but-were-afraid-to-ask-part-1
3. https://stackoverflow.com/questions/54360408/docker-container-udp-communication-with-other-hosts

## Card Swipe Access Matrix

| Card     | Access | FirstCard | 
|----------|--------|-----------|
| 10058397 | N      | Y         |
| 10058398 | N      | -         |
| 10058399 | Y      | -         |
| 10058400 | Y      | Y         |


| Mode            | AntiPassback | FirstCard        | Card     | Granted | FirstCard | Granted | Reason             | Door     |
|-----------------|--------------|------------------|----------|---------|-----------|---------|--------------------|----------|
| Controlled      | N            | none             | 10058398 | N       | N         | N       | 6  no privilege    | Locked   |
| Controlled      | N            | none             | 10058399 | Y       | N         | Y       | 1  swipe ok        | Locked   |
| Controlled      | N            | none             | 10058400 | Y       | Y         | Y       | 1  swipe ok        | Locked   |
| Controlled      | N            | none             | 10058397 | N       | Y         | N       | 6  no privilege    | Locked   |
|                 |              |                  |          |         |           |         |                    |          |
| Normally Open   | N            | none             | 10058398 | N       | N         | N       | 6  no privilege    | Unlocked |
| Normally Open   | N            | none             | 10058399 | Y       | N         | Y       | 1  swipe ok        | Unlocked |
| Normally Open   | N            | none             | 10058400 | Y       | Y         | Y       | 1  swipe ok        | Unlocked |
| Normally Open   | N            | none             | 10058397 | N       | Y         | N       | 6  no privilege    | Unlocked |
|                 |              |                  |          |         |           |         |                    |          |
| Normally Closed | N            | none             | 10058398 | N       | N         | N       | 6  no privilege    | Locked   |
| Normally Closed | N            | none             | 10058399 | Y       | N         | N       | 11 normally closed | Locked   |
| Normally Closed | N            | none             | 10058400 | Y       | Y         | N       | 11 normally closed | Locked   |
| Normally Closed | N            | none             | 10058397 | N       | Y         | N       | 6  no privilege    | Locked   |
|                 |              |                  |          |         |           |         |                    |          |
| Controlled      | N            | firstcard only   | 10058398 | N       | N         | N       | 6  no privilege    | Locked   |
| Controlled      | N            | firstcard only   | 10058399 | Y       | N         | N       | 10 first card      | Locked   |
| Controlled      | N            | firstcard only   | 10058400 | Y       | Y         | Y       | 1  swipe ok        | Unlocked |
| Controlled      | N            | firstcard only   | 10058397 | N       | Y         | N       | 6  no privilege    | Locked   |
|                 |              |                  |          |         |           |         |                    |          |
| Normally Open   | N            | firstcard only   | 10058398 | N       | N         |         |                    |          |
| Normally Open   | N            | firstcard only   | 10058399 | Y       | N         |         |                    |          |
| Normally Open   | N            | firstcard only   | 10058400 | Y       | Y         |         |                    |          |
| Normally Open   | N            | firstcard only   | 10058397 | N       | Y         |         |                    |          |
|                 |              |                  |          |         |           |         |                    |          |
| Normally Closed | N            | firstcard only   | 10058398 | N       | N         |         |                    |          |
| Normally Closed | N            | firstcard only   | 10058399 | Y       | N         |         |                    |          |
| Normally Closed | N            | firstcard only   | 10058400 | Y       | Y         |         |                    |          |
| Normally Closed | N            | firstcard only   | 10058397 | N       | Y         |         |                    |          |
|                 |              |                  |          |         |           |         |                    |          |
| Controlled      | N            | firstcard swiped | 10058398 | N       | N         | N       | 6  no privilege    | Unlocked |
| Controlled      | N            | firstcard swiped | 10058399 | Y       | N         | Y       | 1  swipe ok        | Unlocked |
| Controlled      | N            | firstcard swiped | 10058400 | Y       | Y         | Y       | 1  swipe ok        | Unlocked |
| Controlled      | N            | firstcard swiped | 10058397 | N       | Y         | N       | 6  no privilege    | Unlocked |
|                 |              |                  |          |         |           |         |                    |          |
| Normally Open   | N            | firstcard swiped | 10058398 | N       | N         |         |                    |          |
| Normally Open   | N            | firstcard swiped | 10058399 | Y       | N         |         |                    |          |
| Normally Open   | N            | firstcard swiped | 10058400 | Y       | Y         |         |                    |          |
| Normally Open   | N            | firstcard swiped | 10058397 | N       | Y         |         |                    |          |
|                 |              |                  |          |         |           |         |                    |          |
| Normally Closed | N            | firstcard swiped | 10058398 | N       | N         |         |                    |          |
| Normally Closed | N            | firstcard swiped | 10058399 | Y       | N         |         |                    |          |
| Normally Closed | N            | firstcard swiped | 10058400 | Y       | Y         |         |                    |          |
| Normally Closed | N            | firstcard swiped | 10058397 | N       | Y         |         |                    |          |
|                 |              |                  |          |         |           |         |                    |          |
 

Notes:
1. Door mode _firstcard-only_ is reported as _controlled_.
2. Door mode is set by the most recent _set-firstcard_ or _set-door_ command.
3. Firstcard swipe requires firstcard-only mode.

