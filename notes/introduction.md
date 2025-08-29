# INTRODUCTION TO NETWORKING AND SECURITY TERMS

These are some common terms in security

## HASHING
 - creating fix byte representation of a data
 - commonly used to make sure the source data is the correct one by computing the hash first, and send the hash
   to be used as value to be compared later on
 - common hash algorithm : md4, md5, sha1, sha2, sha3
 - deterministic : same input -> same output
 - characteristic of hashing
	: one way, 
	: fast enough to process gigabytes of data but slow enough to prevent output brute force
	: avoid collision (2 different data has the same output)
	: deterministic