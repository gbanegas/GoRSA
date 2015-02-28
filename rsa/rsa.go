package rsa

import (
        "./primes"
        "math"
        "math/big"
        "math/rand"
        "time"
        "fmt"
       )



func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
}

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

func getUpperLimit(nthPrime int) int64 {
    n := float64(nthPrime)
    x := n * math.Log(n) * 1.2
    return int64(x)
}
 

func GenerateKeys() (*PrivateKey, *PublicKey)  {
    pn := primes.PrimeSieveOfErat(getUpperLimit(999999))
    rand.Seed( time.Now().UTC().UnixNano())
    fmt.Println(pn[randInt(len(pn)/2, len(pn))])
    pi := big.NewInt(pn[randInt(len(pn)/2, len(pn))])
    qi := big.NewInt(pn[randInt(len(pn)/2, len(pn))])
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


