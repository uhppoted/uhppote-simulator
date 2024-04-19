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

2. https://blog.cloudflare.com/everything-you-ever-wanted-to-know-about-udp-sockets-but-were-afraid-to-ask-part-1/

