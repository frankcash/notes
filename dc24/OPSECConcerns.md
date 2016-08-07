# OPSEC Concerns 

Date: 07-Aug-16

# Main

## TL;DR: Patterns and Normacy
- Surveillance does not scale
	- People, malware, packets
- Multiple layers of filtering
- Some Targets are specifically and explicitly tasked

## OPSEC
- Operation Security
	- Keep secrets
	- Not looking like you have secrets
- Compartmentalization
	- Keep services clean
- Signaling vs Communication

## Risk Assessment?
- Who are you protecting yourself from
	- Intelligence services, law enforcement
	- Criminals
	- Comcast

## Don't Think You Are a Target
- US targets tech people
- Not only Gov't to do it

## Why Concerns w/ Crypto
- Crypto is hard
- Encrypting everything is not always the best option in certain circumstances
- Free form text field gives up a lot of information, unique strings etc
	- allows patterns to be drawn on people
- Crypto yells that you have something to hide
	- Metadata
	- Different attack vectors, not everything is covered up completely 
- Commonly get caught on OPSEC

## Problem #1
- Not all encrypted (DNS request, EMAIL headers)
- Even in perfect crypto world, session metadata isn't encrypted

## Key Servers
- Don't die
- People reuse same SSH keys for auth

## Problem #2
- SSL/TLS crts, Signing certs create all sorts of new metadata
	- Geolocation, identity, serial number, create/expiry
- CA's have free form text
- Wildcard certs are dangerous
- Malware builders use malicious builders for certs
- Easy to hunt statistically 
- SSL/TLS is searchable in Shodan
- Re-using certs makes it easy to correlate your activies and break your compartmentalization

## Problem #3
- Encryption can be suspicious
- Targeting abnormality in target selection
### VPNS
- Suspicious
- Size, time, and to who
- Don't even need layer 4
### Make Encryption Mainstream
- If normalized, will be less suspicious 

## What to do
- Threat analysis and the solution

# Tools
- Android-Cert-Generator
	- Problem: Android malware requires you to write fully-functioning app or to trojan another and resign it.  Need a way to create believable.


# Notes

- Ambiguity 
- Follow `@thegruqq`
- Email: john.bambenek@fidelissecurity.com 
