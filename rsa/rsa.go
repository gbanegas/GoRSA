package rsa

import (
        "math/big"
        "crypto/rand"
        "fmt"
       )

var bigOne = big.NewInt(1)

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
    bits := 15
    pi, err := rand.Prime(rand.Reader, bits)
    if err != nil {
            fmt.Println(err)
            return nil,nil
    }
    qi,err := rand.Prime(rand.Reader, bits)
    for {
        if err != nil {
            fmt.Println(err)
            return nil,nil
        }
        if qi.Cmp(pi) != 0 { break }
    }
   
    qe := new(big.Int)
    pe := new(big.Int)
    pe.SetBytes(pi.Bytes())
    qe.SetBytes(qi.Bytes())

    totient := new(big.Int)
    pe.Sub(pe, bigOne)
    qe.Sub(qe, bigOne)
    totient.Mul(pe, qe)
    
    /*ei, err := rand.Int(rand.Reader, ni)
     if err != nil {
        fmt.Println(err)
        return nil,nil ;
    }
    x := new(big.Int)
    y := new(big.Int)
    for {
         result := new(big.Int)
         result.GCD(x, y, qi, ei)
        if result.Cmp(bigOne) == 0 { 
            break 
        }
        ei, err = rand.Int(rand.Reader, ni)
         if err != nil {
            fmt.Println(err)
            return nil,nil
        }
    }*/

    ed := big.NewInt(65537)
    di, ok := modInverse(ed, totient)
    if !ok {
        return nil, nil
    }

    fmt.Println("e = ", ed)
    fmt.Println("pi = ", pi)
    fmt.Println("d = ", di)
    fmt.Println("tot = ", totient)
    return &PrivateKey{p: pi, q: qi, d : di, n : totient}, &PublicKey{n : totient, e : ed}
    
}

func modInverse(a, n *big.Int) (ia *big.Int, ok bool) {
         g := new(big.Int)
         x := new(big.Int)
         y := new(big.Int)
         g.GCD(x, y, a, n)
         if g.Cmp(bigOne) != 0 {
             // In this case, a and n aren't coprime and we cannot calculate
             // the inverse. This happens because the values of n are nearly
             // prime (being the product of two primes) rather than truly
             // prime.
             return
         }
     
         if x.Cmp(bigOne) < 0 {
             // 0 is not the multiplicative inverse of any element so, if x
             // < 1, then x is negative.
             x.Add(x, n)
         }
     
         return x, true
     }  


/*
    c = m^e  mod n
*/
func Cipher(pub *PublicKey, text []byte) ([]byte) {
    c := new(big.Int)
    temp := new(big.Int) 
    temp.SetBytes(text)   
    fmt.Println("D ", pub.e)
    fmt.Println("N ", pub.n)
    fmt.Println("message big int: ", temp)
    c.Exp(temp, pub.e, pub.n)
    
    /*fmt.Printf("%x ", message[i])*/
    return c.Bytes()
}

/*
    m = c^d mod n
*/
func Decipher(priv *PrivateKey, cipheredText []byte) ([]byte) {
    c := new(big.Int)
    m := new(big.Int)
    c.SetBytes(cipheredText)
    fmt.Println("D ", priv.d)
    fmt.Println("N ", priv.n)
    m.Exp(c, priv.d, priv.n)
   
    /*fmt.Printf("%x ", message[i])*/
    
    fmt.Printf("\n")
    return m.Bytes()
}


