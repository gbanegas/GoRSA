package rsa

import "math/big"

type PrivateKey struct {
    p *big.Int
    q *big.Int
    d *big.Int

}

type PublicKey struct {
    n *big.Int
    e *big.Int
}

func GenerateKeys() (*PrivateKey, *PublicKey)  {
    pi := big.NewInt(0)
    qi := big.NewInt(0)
    di := big.NewInt(0)
    ni := pi.Mul(pi, qi)
    ei := big.NewInt(0)
    return &PrivateKey{p: pi, q: qi, d : di}, &PublicKey{n : ni, e : ei}
    
}


/*
    c = m^e  mod n
*/
func Cipher(pub PublicKey, text []byte) {

}

/*
    m = c^d mod n
*/
func Decipher(priv PrivateKey, cipheredText []byte) {

}


