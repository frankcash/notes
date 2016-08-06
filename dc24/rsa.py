from cryptography.hazmat.backends import default_backend
from cryptography.hazmat.primitives import hashes
from cryptography.hazmat.primitives.asymmetric import padding
from cryptography.hazmat.primitives.asymmetric import rsa

# Generate RSA
private_key = rsa.generate_private_key( public_exponent=65537, key_size=2048, backend=default_backend())

# to get the public key
public_key = private_key.public_key()

# encrypt w/ RSA and padding

message = b"The Secret Key"
ciphertext = public_key.encrypt(message, padding.OAEP(mgf=padding.MGF1(algorithm= hashes.SHA1()), algorithm=hashes.SHA1(), label=None))
plaintext = private_key.decrypt(ciphertext, padding.OAEP(mgf=padding.MGF1(algorithm = hashes.SHA1()), algorithm=hashes.SHA1(), label=None))
