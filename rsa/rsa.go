package rsa

import (
        "math/big"
        "fmt"
       )

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
    pe := pi.Sub(pi, big.NewInt(1))
    qe := qi.Sub(qi, big.NewInt(1))
    fmt.Println("pe = ", pe)
    fmt.Println("qe = ", qe)
    fmt.Println("pi = ", pi)
    fmt.Println("qi = ", qi)
    totient := pi.Mul(pe, qe)
    fmt.Println("tot = ", totient)
    ni := pi.Mul(pi, qi)
    ei := big.NewInt(0)
    fmt.Println("pi = ", pi)
    return &PrivateKey{p: pi, q: qi, d : di}, &PublicKey{n : ni, e : ei}
    
}


/*
    c = m^e  mod n
*/
func Cipher(pub *PublicKey, text []byte) (cipheredText []byte) {
    message := ""
    temp := big.NewInt(0)
    m := temp.SetBytes(text);
    m = m.Exp(m,pub.e, pub.n)
    message = message + m.String()
    fmt.Println(message)
    /*fmt.Printf("%x ", message[i])*/
    
    fmt.Printf("\n")
    return nil
}

/*
    m = c^d mod n
*/
func Decipher(priv PrivateKey, cipheredText []byte) {

}


