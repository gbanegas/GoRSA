package rsa

import (
        "math/big"
        "crypto/rand"
        "io"
        //"fmt"
        //"time"
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

var smallPrimes = []uint8{
    		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
    }
var smallPrimesProduct = new(big.Int)


func GenerateKeys(bits int) (*PrivateKey, *PublicKey)  {
    //fmt.Println("Generating keys....... with the size of : ", bits)
    //rand.Seed(time.Now().Unix())
    smallPrimesProduct.SetString("2305567963945518424753102147331756070",10)
    pi, err := prime(rand.Reader, bits)
    if err != nil {
            //fmt.Println(err)
            return nil,nil
    }
    qi,err := rand.Prime(rand.Reader, bits)
    for {
        if err != nil {
            //fmt.Println(err)
            return nil,nil
        }
        if qi.Cmp(pi) != 0 { break }
    }
   
    qe := new(big.Int)
    pe := new(big.Int)
    pe.SetBytes(pi.Bytes())
    qe.SetBytes(qi.Bytes())
    pe.Sub(pe, bigOne)
    qe.Sub(qe, bigOne)

    ne := new(big.Int)
    ne.Mul(pi,qi)
    
    totient := new(big.Int)
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
    return &PrivateKey{p: pi, q: qi, d : di, n : ne}, &PublicKey{n : ne, e : ed}
    
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
func Cipher(pub *PublicKey, m []byte) ([]byte) {
    c := new(big.Int)
    temp := new(big.Int) 
    temp.SetBytes(m)   
    c.Exp(temp, pub.e, pub.n)
    return c.Bytes()
}

/*
    m = c^d mod n
*/
func Decipher(priv *PrivateKey, cipheredText []byte) ([]byte) {
    c := new(big.Int)
    m := new(big.Int)
    c.SetBytes(cipheredText)
    m.Exp(c, priv.d, priv.n)
    return m.Bytes()
}


func prime(rand io.Reader, bits int) (p *big.Int, err error) {
 			b := uint(bits % 8)
    		if b == 0 {
    			b = 8
    		}
    	
    		bytes := make([]byte, (bits+7)/8)
    		p = new(big.Int)
    	
    		bigMod := new(big.Int)
    
    		for {
    			_, err = io.ReadFull(rand, bytes)
    			if err != nil {
    				return nil, err
    			}
    	

    		bytes[0] &= uint8(int(1<<b) - 1)

    			if b >= 2 {
    				bytes[0] |= 3 << (b - 2)
    			} else {
    				// Here b==1, because b cannot be zero.
    				bytes[0] |= 1
    				if len(bytes) > 1 {
    					bytes[1] |= 0x80
    				}
    			}

    			bytes[len(bytes)-1] |= 1
    	
    			p.SetBytes(bytes)
    	
    			bigMod.Mod(p, smallPrimesProduct)
    			mod := bigMod.Uint64()
    	
    		NextDelta:
    			for delta := uint64(0); delta < 1<<20; delta += 2 {
    				m := mod + delta
    				for _, prime := range smallPrimes {
    					if m%uint64(prime) == 0 && (bits > 6 || m != uint64(prime)) {
    						continue NextDelta
    					}
    				}
    	
    				if delta > 0 {
    					bigMod.SetUint64(delta)
    					p.Add(p, bigMod)
    				}
    				break
    			}
    	
      			if p.ProbablyPrime(20) && p.BitLen() == bits {
   				return
   			}
   		}
}


