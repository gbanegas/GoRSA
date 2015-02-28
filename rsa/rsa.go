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

func GeneratePrivateKey() (*PrivateKey)  {
    pi := big.NewInt(0)
    qi := big.NewInt(0)
    di := big.NewInt(0)
    return &PrivateKey{p: pi, q: qi, d : di}
    
}
func GeneratePublicKey(priv *PrivateKey) (*PublicKey){
    ni := priv.p.Mul(priv.p, priv.q)
    ei := big.NewInt(0)
    return &PublicKey{n : ni, e : ei}
}

func Cipher(pub PublicKey, s []byte) {

}


