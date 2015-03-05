package main

import (
        "./rsa"
        "math/big"
        "fmt"
       )

func main(){
    fmt.Println("RSA Implementation")
    bits := 1024
    priv , pub := rsa.GenerateKeys(bits)

    fmt.Printf("private key : %#v\n", priv)
    //fmt.Printf("%#v\n", priv)
    //fmt.Printf("%T\n", priv)
    //fmt.Printf("%p\n", &priv)
    fmt.Printf("public key : %#v\n", pub)

    sample := big.NewInt(1234)
    fmt.Println("Original plain Text : ", sample)

    //To encrypt
    ciphered := rsa.Cipher(pub,sample.Bytes())
    ciphered_text := new(big.Int)
    ciphered_text.SetBytes(ciphered)
    fmt.Println("ciphered text = ", ciphered_text)

    //To decrypt
    result := rsa.Decipher(priv, ciphered)
    c := new(big.Int)
    c.SetBytes(result)
    fmt.Println("plain text = ", c)

}