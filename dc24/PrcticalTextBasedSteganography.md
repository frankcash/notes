# Practical Text-Based Steganography: Exfiltrating data from secure networks and socially engineering SecOps Analysts

By: TryCatchHCF
Date: August 05, 2016

# Main
- Data Loss Prevention (DLP). Know what to look out for.
- Secure networks can be defeated by lists of strings: cookies, etc
- "Vurt da Furk!"

## Data Exfiltration 
- Getting data out of network is better than just in
Cats of Exfiltration
- Standard data transfer
- Out-of-band
- Obscure channels
- Physical

- Image Steganography is becoming more contrived


## Advantages of standard data approach
- Does not depend on presence of port, protocol, or app
- Allows flexibility on part of attacker
- Avoid having to infiltrate and install tools
- No need to infiltrate physical devices onto target subnet

## Advantages to text-based steganography
- text is a universal format

## Cloakify Toolset
- [Github](github.com/TryCatchHCF/Cloakify)
- `cloakify.py` - obscure data prior to exfiltration
- `decloakify.py` - Decode exfil data
- Very simple
- First script base64 encodes => applies a cipher to generate a list of strings representing characters in the b64 payload

### Not a secure encryption scheme
- Prone to freq analysis
- Encrypt data separately prior to cloaking to keep secure
- Python 2.7.x
- More so encoding instead of encryption

## Freq Analysis Vuln
- Since at most 66 strings are used, computationally easy for machines to detect when ciphers are in use
- Can add entrophy. I.e. through prepending things such as sequential timestamps. 

## Create your own cipher
- Generate a list of at least 66 unique words/ phrases/ symbols
- Randomize the list order
- Remove all duplicate entries and all blank lines

# Notes
- Accepting pull requests
