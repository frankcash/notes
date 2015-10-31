Title: Cameras, Thermostats, and Home Automation Controllers
Subtitle: Hacking 14 IoT Devices
Time: Friday, Aug 7, 1200

(Initial testing was in November 2014)


Over View of Testing
- Goals: Figure out what we want to get out
- "What is someone likely to do to make us look bad" generally wont care about security after that
- Three diff ways for IoT: Physical, local network (can mean BLE or LAN), remote
- Physical < LAN < Remote
- Hardware ? PC Apps? Mobile Apps? Cloud Components? - Don't limit yourself

Techniques
- Hardware: expensive equip. not needed (multimeter, serial cable, etc)
- PC Apps: monitor communications (mitm?), determine how trust is established, check for poor development practices
- Mobile Apps: monitor (MITM?), decompile app, determine trust, general bad app dev practices
- Cloud Comms: Auth, general security, initial registration

Results:
- DLink DCS-2132L: Custom cleartext UDP, uses plain HTTP, (pc) backups contain cleartext passwords, (pc) auto updates from plain HTTP
- DropCam Pro (generally secure): exposed serial port, built on top Ambarella dev kit
- FOSCAM FI9826W (history of vults): Plaintext by default (port 88), backs encrypted to static key, (Mobile) hardcoded API keys, data written to SD instead of cloud
- Closeli Simplicam (some fixed): /hom/.config contains root passw, firmware updates encrypted, doesn't verify certs, init connect over http, (Mobile) hardcoded API Keys, (Mobile) sessions and what not stored in phone filesyste, (cloud) XMPP over HTTP for vid streams
- Withings Baby Monitor: SSH part with fixed root, HTTP, tons of GPL code (not released), (mobile) GCM key exposed in app
- Ecobee (generally secure): has serial port, HTTPS, no cert pinning, torx t6 -> use pins -> boots to root shell
- Hive (generally secure): mix of HTTP/HTTPs, no cert pinning, (cloud) on device references to test instances, boot -> user root -> password <blank>
- Honeywell Lyric (somewhat secure): runs web server during initial pairing, pairing w/ HTTP, not linux, (mobile) creds can show up in logs when crash, (cloud) uses static IPs 
- Nest Thermostat (quite secure): Linux, can be rooted via usb, (Mobile) uses HTTPS, (cloud) weather uses HTTP not HTTPS, htvhacker (exploiteers) put out roots
- Nest Protect (yet to be rooted/ secure right now): during pairing acts as open WiFi service (mess with this)
- Control4 HC-250: No network controlls, default root passw, only decent if firewalled, filesystem stored in SD card, has webserv
- Lowes Iris (quite secure): serial shell is easy
- Revolv: support reverse SSH tunnel, firmware is amazing, 6+ debug / jtag ports, lots of open TCP ports, (mobile) device creds stored in sqlite db, (mobile) some API keys secured , uses custom encrypt from java (lol wtf)


slides: syn.ac/defcon2015iot
