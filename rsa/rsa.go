package rsa

import (
        "math/big"
        "crypto/rand"
        "fmt"
       )





type PrivateKey struct {
    p *big.Int
    q *big.Int
    d *big.Int
    n *big.Int

}


type PublicKey struct {
    n *big.Int
    e *big.Int
}

 

func GenerateKeys() (*PrivateKey, *PublicKey)  {
    pi, err := rand.Prime(rand.Reader, 99)

   if err != nil {
      fmt.Println(err)
   }
    
    
    qi := big.NewInt(0)
    di := big.NewInt(0)
    qe := big.NewInt(0)
    pe := big.NewInt(0)

    pe.SetBytes(pi.Bytes())
    qe.SetBytes(qi.Bytes())

    totient := big.NewInt(0)
    pe.Sub(pe, big.NewInt(1))
    qe.Sub(qe, big.NewInt(1))
    fmt.Println("pe = ", pe)
    fmt.Println("qe = ", qe)
    fmt.Println("pi = ", pi)
    fmt.Println("qi = ", qi)
    totient.Mul(pe, qe)
    fmt.Println("tot = ", totient)
    ni := pi.Mul(pi, qi)
    ei := big.NewInt(0)

    fmt.Println("pi = ", pi)
    return &PrivateKey{p: pi, q: qi, d : di,n : ni}, &PublicKey{n : ni, e : ei}
    
}


/*
    c = m^e  mod n
*/
func Cipher(pub *PublicKey, text []byte) (string) {
    message := ""
    temp := big.NewInt(0)
    m := temp.SetBytes(text);
    m = m.Exp(m,pub.e, pub.n)
    message = message + m.String()
    fmt.Println(message)
    /*fmt.Printf("%x ", message[i])*/
    
    fmt.Printf("\n")
    return message
}

/*
    m = c^d mod n
*/
func Decipher(priv PrivateKey, cipheredText []byte) (string) {
    message := ""
    temp := big.NewInt(0)
    c := temp.SetBytes(cipheredText);
    c = c.Exp(c, priv.d, priv.n)
    message = message + c.String()
    fmt.Println(message)
    /*fmt.Printf("%x ", message[i])*/
    
    fmt.Printf("\n")
    return message
}


