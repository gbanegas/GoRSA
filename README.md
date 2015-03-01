# GoRSA
RSA implemented in Go language. 

A simple implemantion of RSA in Go. 

# Operation

## Encryption

m - message (plain text)
(e, n) - public key
c - ciphered Text

c = m^e (mod n)

## Decryption

c - ciphered text 
(p, q, d) - private key

m = c^d (mod n)



