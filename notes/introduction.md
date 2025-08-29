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
   - DSA -> secure-ish
   - RSA
   - Diffie-Hellman
   - ECDSA
   - ECDH

## DIGITAL SIGNATURE
- A technique to make sure the data or message is not tempered( integrity of the data is not compromised) and to make sure
  it is the appropriate user sends it (authenticity).
- This can be achieved by using the asymmetric encryption technique and hashing.
- Considered this case where a sender want to send `A` a file to a client `B`. 
- `A` can send data directly to `B` but the problem is what happen if there is some interruption in the middle. It can 
  alter the value that `A`sends. So it can be achieved by encrypting the data by `A`'s private key first. And then
  send it. If `B`, who has the public key of `A`, can decrypt it successfuly, it can be guaranteed that it is the key pair
  that `A` issued earlier. But the problem is, that the data itself can be bulky so it doesn't really a good choice to 
  encrypt the actual message using asymmetric encryption. So this is where the `hash` comes into play. Hashing algorithm
  can produce digital 'fingerprint' of the message with fixed-length. After hashing it, `A` can encrypt this hashed string
  using its private key and further sends it to `B`. The `B` can decrypt it using public key and compute the message hash.
  If the received hash is the same with the computed hash, then it can be guaranteed that the data sent is not compromised.
```
A → Compute hash → Sign hash with private key → Send message + signature.
B → Compute hash of received message → Verify signature using public key.
```
- But remember that this only covers the integrity and authenticity of a message. Because the message is sent without
  encryption. That is why this approach is not for sending confidential information but only to prove that a message is
  not tempered and issued by the valid sender.
- The application can be used to send certificate or some message. This will be covered for TLS protocol later on.

## Cryptography and Web Security 
- Web security can be achieved by incorporates cryptography techniques like hashing, encryption, and signature.
- TLS, stands for Transport Layer Security, uses this technique to ensure integrity, authenticity and confidentiality of 
  web communication.
- It exist between the HTTP layer and Transport Layer on the network stack.
- To make data is sent securely, it can be achieved by encrypt it using symmetric key encryption. But remember previous 
  explanation where it is considered less secure because whenever the key is compromised, its done. But if you use
  asymmetric key instead, it is considerably slow and bulky especially for bigger message. So the best approach is to 
  use `Hybrid Approach`, that is using both symmetric and asymmetric key for encryption
- Encryption of the actual message is done by the symmetric ones. But how to make sure it can't be accessed by attacker?
  It can be achieved by encrypting the key itself, not the message. So the server sends the client its own public key.
  For encryption, you use public key, right? So the client, who receives the public key, encrypt the symmetric key,
  and sends it to the server. The server then decrypt it using private key which is only held by the server. The server 
  now has the symmetric key and can send the data in encrypted format without worrying of spies. The client originally hold
  the symmetric key too, so it can decrypt it. No other parties that can understand the data sent because they dont have
  any symmetric key.
- That is how the internet send the data securely, but.... there is caveat. How does the client know that the servers' public
  key sent by the server is the authorized key? There is a chance that another party intercepts the communication, acts
  like a legit server and sends its own public key to the client. The process would be run smoothly as previously
  explained. But the problem is, the client didn't know the symmetric key exchange is not from a legit server.
  So this is where the `Certificates Authority` or CA comes in.
- CA is a certification issued by a several trusted parties/companies that give a validation or signature to ensure
  that the key is indeed owned by the authorized server.
- This parties/companies must be trusted by client.
- So the server request the CA to sign its certification which contains the domain info, public key and any related data. 
  When the client receives it, it can see that the key is not tampered and own by the actual server.
- So the CA acts like 'This public key is indeed issued by the example.com. So you can use it safely'. 
