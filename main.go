package main

import (
        "./rsa"
        "math/big"
        "fmt"
       )

func main(){
    fmt.Println("Hello, playground")
    priv , pub := rsa.GenerateKeys()

    fmt.Println("priv = ", priv)
    fmt.Println("pub = ", pub)
    sample := big.NewInt(1234)

    ciphered := rsa.Cipher(pub,sample.Bytes())
    fmt.Println(ciphered)
    result := rsa.Decipher(priv, ciphered)
    c := new(big.Int)
    c.SetBytes(result)
    fmt.Println("result = ", c)

}