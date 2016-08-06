# Getting Started With Crypto in Python

Date: August 05, 2016


## Github

github.com/amiralis/python_crypto_cpv

# Main

- Crypto is ubiquitos

pycrypto (oldest)
m2crypto (swig binding)
[cryptography (PY2, PY3)](https://cryptography.io/en/latest/)

`pip install cryptography`

## Cryptography.io
- Primitive Crypto Blocks (cryptography.hazmat)
	- Message digest and hashing algos
	- Symmetric encryption algos
	- Asymmetric encryption algos
- X.509 Ecosystem
- Full high level crypto recipe

## Hashing Algorithims
- Input: Long message
- Output: short black
- Collisions: It is computationally difficult to find two messages such that h(m) = h(m^2)
- Examples:
	- Recommended Hash Algos (SHA-2, SHA-3) by NIST
	- SHA-1: output 160bits; being phased out
	- MD2, MD4, MD5 by Ron Rivest [RFC1319, RFC1320, RFC1321]

### SHA Family
- Secure Hash Algorithm Family.
	- Just use Sha 3
```
from cryptography.hazmat.backends import default_backend
from cryptography.hazmat.primitives import hashes
import base64

for _hash in [hashes.SHA1, hashes.SHA224, hashes.SHA256, hashes.SHA384, hashes,SHA512]:
	digest = hashes.Hash(_hash(), backend=default_backend())
	digest.update(b"CryptoVillage")
	digest.update(b"2016"),
	msg_digest = digest.finalize()
	print _hash.name, len(msg_digest), len(msg_digest) * 8
```

### Hash-based message authentication code (HMAC)
- HMACS are used for messaging authentications combined w/ a secret key. They provide integrity check and auth

### RSA
- RSA, is an asymmetric encryption algorithm
- Security is based on factorization problem
- RSA problem: it's slow and not used for encrypting large data
- Most used to encrypt the symmetric key that is used for encryption

### OpenSSL
Can generate keys

```bash
# generate a 2048 bit private key
openssl genrsa -out private_key.pem 2048
openssl pkcs8 -topk8 -inform PEM -outform DER -in private_key.pem -out private_key.pem
openssl rsa -in private_key.pem -pubout -outform DER -out public_key.der
```
```python
# generate a 2048 bit private key
from cryptography.hazmat.backends import default_backend
from cryptography.hazmat.primitives.asymmetric import rsa

private_key = rsa.generate_private_key( public_exponent=65537, key_size=2048, backend=default_backend())

# to get the public key
public_key = private_key.public_key()

```
```python
bin(2**16 + 1) // is an easy computation because the math is faster
```
- Textbook RSA is not IND-CPA secure, therefore we need to use Optimal Asymmetric Encryption Padding (OAEP).
- It's all about padding

```python
from cryptography.hazmat.backends import default_backend
from cryptography.hazmat.primitives import hashes
from cryptography.hazmat.primitives.asymmetric import padding
message = b"The Secret Key"
ciphertext = public_key.encrypt(message, padding.OAEP(mgf=padding.MGF1(algorithm= hashes.SHA1()), algorithm=hashes.SHA1(), label=None))
plaintext = private_key.decrypt(ciphertext, padding.OAEP(mgf=padding.MGF1(algorithm = hashes.SHA1()), algorithm=hashes.SHA1(), label=None))

```

## Symmetric Encryption

### AES - Andvanced Encryption Algorithm
- has different modes of operations
#### Block cipher mode of operation
- to encrypt message of arbitrary size w/ block ciphers we use following algos.  They define how to encrypt each block of the plaintext to produce corresponding cipher text block. ECB is extremely insecure.
	- Electronic Codebook (ECB) Not secure. 
	- Cipher Block Chaining (CBC) Somewhat secure. Needs an IV.
	- Counter (CTR) Somewhat secure. Pretty fast.  Does not use passing.

- Snacphat was using ECB (PWNED)

## Pretty Good Privacy (PGP)
- created by Phil Zimmerman in 1991
- Uses both symmetric and asymmetric crypto

# Notes
`pip install jasmin`
`pip install cryptography`
