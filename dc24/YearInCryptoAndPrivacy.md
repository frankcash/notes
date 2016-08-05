# This Year in Crypto & Privacy
By: Organizers of Crypto Village
Date: August 05, 2016

# Main

## LabMD
- FTC does consumer protection
- FTC Alleges company failed to reasonably protect security of personal data including med info
- Two Seperate Incidents: Exposure of PII ~ 10k Consumers
	- 9k Consumers were found on a p2p network
	- LabMD docs containing sensitive info of 500 consumers in hands of identity thievs
- LabMD challenged the FTC's Authority 
- ALJ dismisses complaint: FTC "failed to carry its burden of proving its theory that Respondent's failure to employ reasonable data security constitutes an unfair trade practice because they didn't prove part of three prong test"
- Got appealed to comissioners: ALJ applied wrong test. LabMD failed to use intrusion detection, lacked basic precautions to protect the sensitive consumer

## Wyndham
- FTC sued for not protecting data security.  (3 data breaches in 2 yrs).
- District court said: FTC had power to do such, Wyndham failed to have proper data security.  Circuit court agreed.

## Ethereum 
- Public blockchain crypto currency
- Built in VM; turing complete; smart contracts; prediction market
- Some company using it raised 150m; raising the price of ether
- This company (DAO) had a vuln.
- Hacker found way to by-pass voting process for project decision, funneled ether to his $$.
- Ethereum company had a hard fork of block chain. 
- Ether went to $10.  Hard fork went through.
- Consider further auditing of smart contracts.  Turing complete + Smart Contracts = fucking hard.

## Crypto Coming to Mobile
- Not all crypto is equal
- Use Tor. Use Signal (Free to audit) #opsec.
### Why signal?
- Open Source
- Not everyone uses Signal
	- Lot of other companies can license Signal tech (i.e. WhatsApp).
	- WhatsApp backs up data, which is not encrypted.
	- WhatsApp doesn't delete shit (see cypherpunks).
	- Google's Allo (coming soon). Have to enable settings to incognito in settings
	- FB Messenger (coming soon). Will have to manually set
- ripperino Tox :(

## Changes to TLS
- Protocol that encrypts the web
- SSL => TLS
- Tons of named vulns in past 5
### Top HTTPS Adoption Issues
- Cert Management
- Mixed Content
- added Latency
	- TLS 1.3 wants to make this fast and clean up broken crypto
	- 1.3 removes one handshake. Removed some older crypto (RSA static mode, etc)
	- 1.3 pls get standarized.  Draft 14. Implementations: BoringSSL (chrome), Mint, tls-tris, NSS (Firefox)

## Year in "Secure Backdoors"
- The NSA's secure backdoor
- Encryption is freely available
- Many crypto schemes break if you can predict random numbers
- NSA proposed a abd random number gen as a NIST standard (DUAL_EC_DRBG: Adopted 2006) Could've proved it was backdooored.
- Great backdoor. relied on PKI crypto that only the NSA could predict.
Christmas 2015
- Juniper Security Advisory: a knowledge attacker could monitor VPN traffic and decrypt (CVE-2015-7758). 
- Internet begins diffing firmware binaries
- Juniper used Dual_EC_DRBG.  Does scrambling to avoid state logging.
- Was using global var, never ran rnadom numb genertor to get raw output of backdoored number gen
- At same time changed IPsec key setup to 32bytes which allowed perfect recover of Dual_EC_DRBG state
- Someone changed magic param to intercept VPN traffic (ScreenOS)
- Secure backdoors for exceptional, lawful access are a joke.

## Breaking RSA followup
- NSA announced they're planning on transition to quantum resistant algorithms. (2015)
- In the 1990s a bunch of quantum attacks were announced
- NSA => Suite B to Quantum Resistant Algos
- Will be using a challenge through NIST until 2020

## Apple vs. FBI
- iPhone 5c was used by shooter (has passcode)
- FBI was able to access iCloud
- Weeks before stopped syncing w/ iCloud
- All Writs Act, 1789 (very old)
- Tim Cook refused creating backdoor. Citing dangerous precedent. 

## Encryption Based on Information Theory
- Encryption and decryption NOT based solely on key space
	- Creating group of keys giving the same output
	- Easier to break things
- Polymorphism in Cryptography
	- mutating cyphers
	- Small submessages
- Side channel attacks on modes
	- possible to break CBC and ignore the message
	- Death to CBC soon
	- Gonna have to stop using big keys
- CBC broken for collisions, non collision attacks being researched

## Govt Interference in Cryptography
- Companies using laws to get leg up on competition, prevents better algos
- Kazakhstan Security Cert
- France outlawing Anonymous

## HTTPS for all w/ Let's Encrypt
- Went live in December
- Tools came out to make this process easy
- ~5mins to get cert
- WordPress, blogr, etc using this
Safe and secure
- 6 months - 7 million domains have Let's Encrypt Cert.

## Changing passwords frequently is not more secure :^)
- GCHQ: stop changing frequently
- Has been researched, when you enforce mandatory passwd changes actually drops burdens and increases usage of weak passwords
