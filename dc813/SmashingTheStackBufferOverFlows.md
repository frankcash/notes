Title: Smashing the Stack Old School!
Date: August 19, 2015

Nehalem: https://en.wikipedia.org/wiki/Nehalem_%28microarchitecture%29

Pre-Nehalem:
-No randomization 

Idea: pour in more bytes, then manipulate return address to point to shell code (controlled crash)

Stack frames designate beginning and end function

Step 1: run program to buffer overflow
Step 2: find any open sockets, etc
Step 3: send buffer
Step 4: find where buffer is overflowing
Step 5: Have point to shell code

Prevention: GNU Stack Elf
