# Evil ESP

By: Matt Trimble & Eric Escobar
Date: August 06, 2016

# Main

"Ghetto but gets the job done"

Can be used to find APs

Does not support full promiscuous mode

## ESP 8266

- Microcontroller attached to a WiFi chip
- 32bit RISC CPU @ 80mhz
- 512kb flash
- WiFi supports 802.11/g/n in 2.4GHz freq
- UART
- Pins may be accessible depending on model
- GCC support
- Arduino IDE support
- WiFi chips are FCC certified

## Tools

1. Get Arduino IDE
2. [Add arduino esp8266 to additional boards](arduino.esp8266.com/stable/package_esp8622com_index.json)
3. Add ESP board via Tools -> Board:[XXX] -> 
4. Make programmer
	- not breadboard compatible
	- need serial to usb w/ 3.3v pin
	- will be jank af
5. First Sketch
	- Add include file "ESP8266WiFi.h"
	- github.com/esp8266/arduino

# Notes

[slides](ragingsecrity.ninja)
