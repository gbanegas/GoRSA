# GoRSA
RSA implemented in Go language. 

A simple implemantion of RSA in Go. 

# Operation

## Key Generation

    Choose two distinct prime numbers p and q.
        For security purposes, the integers p and q should be chosen at random, and should be of similar bit-length. Prime integers can be efficiently found using a primality test.
    Compute n = pq.
        n is used as the modulus for both the public and private keys. Its length, usually expressed in bits, is the key length.
    Compute φ(n) = φ(p)φ(q) = (p − 1)(q − 1) = n - (p + q -1), where φ is Euler's totient function.
    Choose an integer e such that 1 < e < φ(n) and gcd(e, φ(n)) = 1; i.e., e and φ(n) are coprime.
        e is released as the public key exponent.
        e having a short bit-length and small Hamming weight results in more efficient encryption – most commonly 216 + 1 = 65,537. However, much smaller values of e (such as 3) have been shown to be less secure in some settings.[5]
    Determine d as d ≡ e−1 (mod φ(n)); i.e., d is the multiplicative inverse of e (modulo φ(n)).

            This is more clearly stated as: solve for d given d⋅e ≡ 1 (mod φ(n))
            This is often computed using the extended Euclidean algorithm. Using the pseudocode in the Modular integers section, inputs a and n correspond to e and φ(n), respectively.
            d is kept as the private key exponent.



## Encryption

m - message (plain text)

(e, n) - public key

c - ciphered Text

c = m^e (mod n)

## Decryption

c - ciphered text 

(p, q, d) - private key

m = c^d (mod n)



