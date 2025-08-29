# INTRODUCTION TO NETWORKING AND SECURITY TERMS

These are some common terms in security

## HASHING
 - creating fix byte representation of a data. it is like a 'fingerprint' of a source data/message.
 - commonly used to make sure the source data is the correct one by computing the hash first, and send the hash
   to be used as value to be compared later on
 - common hash algorithm : md4, md5, sha1, sha2, sha3
 - deterministic : same input -> same output
 - characteristic of hashing
	: one way, 
	: fast enough to process gigabytes of data but slow enough to prevent output brute force
	: avoid collision (2 different data has the same output)
	: deterministic


    https://www.youtube.com/watch?v=aCDgFH1i2B0
## ENCRYPTION
https://www.youtube.com/watch?v=o_g-M7UBqI8
 - Encryption is a technique to ensure a message/data can only be read/understood by a specific users.
   This is achieved by creating it like a random data to a non-targeted users using cryptographic algorithm
 - Encryption can be created using many technique, but for know using key-based encryption.
 - There are two ways: Symmetric and Asymmetric Encryption.

### Symmetric Key
 - The key to encrypt and decrypt is the same.
 - Encrypted output has the same length compared to the original ones.
 - The longer the key, (generally) the better.
 - Because the output has the same length, it is generally used for encrypting a bulk message.
 - Also considered less secure because whenever unauthorized users get the key, it can easily decrypt the message
   because the encryption and decryption is the same key
 - Example of encryption algorithm:
   - DES (56 bit key) -> compromised ❌
   - RC4 (128 bit key) -> compromised ❌
   - 3DES (168 bit key) -> secure-ish
   - AES (128, 192, 256 bit keys) -> recommended ✅
   - ChaCha (128, 256 bit keys) -> recommended✅

### Asymmetric Key
 - The key to encrypt and decrypt is different.
 - Commonly called public key and private key. The public key is the key that can be safely shared to another users.
   Where the private key is the reverse
 - If some encryption is done by the private key, the decryption can only be done by the public key and vice versa.
 - The encryption also produces quite longer messages than the original.
 - This is considered more secure than the symmetric ones due to the asymmetric capability of the key itself
 - But the downside is that the encryption is generally slower and bulkier so it is not generally recommend to use it
   for bulk message/data.
 - Private and public key can be used to both encrypt and decrypt. But you have to know when to use which because 
   it has its own caveat.
 - `Private key: encrypt; Public key: decrypt` => this is for creating sign. To prove that a certain message is issued
   by the holder of the private key. Don't use this to encrypt message that you don't want to be seen by unauthorized 
   because they will have public key and they can decrypt is easily.
 - `Public key: encrypt; Private key: decrypt` = > this is for ensuring confidentiality, because the instance that can
   understand the message, is the only one who has the private key, which is the right target. If you use this backwards,
   then the users can decrypt you private message easily. 
 - Example of encryption algorithm:
   - DSA
   - RSA
   - Diffie-Hellman
   - ECDSA
   - ECDH

