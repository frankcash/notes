# Attacking Hypervisors via Firmware and Hardware

By: Intel Security Advanced Threat Research

# Hypervisor based isolation:
- If a VM part gains privilege a Hypervisor usually protects
- Hypervisor developer nt thinking about hardware => firmware versions when protecting themselves
- A lot of have some device isolation (interrupt remapping, etc)
- (Hypervisor traps are VM exits)

# Firmware RootKit vs Hypervisor
- Firmware rootkit is something affecting System (BIOS or EUFI)

# Installing RootKit
- Physical access
- From privileged gues
- If access to VMM

System doesn't properly protect firmware in SPI flash memory we could bypass write-protection

Malware can exploit vulnerabilities in firmware to install a rootkit on such systems

# VMM Forensics:
- With the help of a rootkit in firmware any VM guest can extract all information about hypervisor and VMs from memory

# Attacking Hypervisors through System Firmware (done from privileged position):
- Pointer vulnerabilityies in SMI Handlers.  Exploit tricks SMI handler to write to an address inside SMRAM. VMM allows VM to invoke SMI Handlers (grants access to SW SMI I/O Port). Compromised VM injects SMM load.  Then takes control of VMM.

# Tools and Mitigation:
- Firmware can be tested for vulns
- Protect firmware in system flash memory
- Chipsec has tests and is open source

# Conclusions:
- Compromised firmware is bad news for VMM
- Make sure privileged guests are hardened because they can easily pwn hypervisor
- Vulns in device and CPU emulation are common. Fuzz all HW interactions

github.com/chipsec
