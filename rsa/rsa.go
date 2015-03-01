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
    bits := 10
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
    fmt.Println("pe = ", pe)
    fmt.Println("qe = ", qe)
    fmt.Println("pi = ", pi)
    fmt.Println("qi = ", qi)

    totient.Mul(pe, qe)
    ni := new(big.Int)

    fmt.Println("tot = ", totient)
    ni = ni.Mul(pi, qi)
    ei, err := rand.Int(rand.Reader, ni)
     if err != nil {
        fmt.Println(err)
        return nil,nil ;
    }
    for {
        result :=  qi.GCD(big.NewInt(1), big.NewInt(1), qi, ei)
        if result.Cmp(bigOne) == 0 { 
            break 
        }
        ei, err = rand.Int(rand.Reader, ni)
         if err != nil {
            fmt.Println(err)
            return nil,nil
        }
    }

    di, ok := modInverse(ei, ni)
    if !ok {
        return nil, nil
    }

    fmt.Println("e = ", ei)
    fmt.Println("pi = ", pi)
    fmt.Println("d = ", di)
    return &PrivateKey{p: pi, q: qi, d : di, n : ni}, &PublicKey{n : ni, e : ei}
    
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
func Cipher(pub *PublicKey, text []byte) (string) {
    message := ""
    temp := big.NewInt(0)
    m := temp.SetBytes(text);
    m = m.Exp(m,pub.e, pub.n)
    message = message + m.String()
    /*fmt.Printf("%x ", message[i])*/
    return message
}

/*
    m = c^d mod n
*/
func Decipher(priv *PrivateKey, cipheredText []byte) (string) {
    message := ""
    temp := big.NewInt(0)
    c := temp.SetBytes(cipheredText);
    c = c.Exp(c, priv.d, priv.n)
    message = message + c.String()
    /*fmt.Printf("%x ", message[i])*/
    
    fmt.Printf("\n")
    return message
}


