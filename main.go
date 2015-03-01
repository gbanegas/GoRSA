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
    sample := big.NewInt(1)

    ciphered := rsa.Cipher(pub,sample.Bytes())

    fmt.Println(ciphered)
    ciphered_bytes := []byte(ciphered)

    result := rsa.Decipher(priv, ciphered_bytes)
    fmt.Println("result = ", result)

}