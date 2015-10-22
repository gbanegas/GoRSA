package main

import (
        "./rsa"
        "math/big"
        "fmt"
        "time"
        "log"
       )

func main(){
    log.Printf("RSA Implementation")
    start_gen := time.Now()
    bits := 8192
    priv , pub := rsa.GenerateKeys(bits)
    elapsed_gen := time.Since(start_gen)
    log.Printf("Key Generation took %s", elapsed_gen)
    fmt.Printf("private key : %#v\n", priv)
    //fmt.Printf("%#v\n", priv)
    fmt.Printf("%T\n", priv)
    //fmt.Printf("%p\n", &priv)
    fmt.Printf("public key : %#v\n", pub)

    sample := big.NewInt(1234)
    //fmt.Println("Original plain Text : ", sample)

    //To encrypt
    start := time.Now()
    ciphered := rsa.Cipher(pub,sample.Bytes())
    ciphered_text := new(big.Int)
    ciphered_text.SetBytes(ciphered)
    elapsed := time.Since(start)
    log.Printf("Encrypt took %s", elapsed)
    fmt.Println("ciphered text = ", ciphered_text)
    //To decrypt
    start_2 := time.Now()
    result := rsa.Decipher(priv, ciphered)
    c := new(big.Int)
    c.SetBytes(result)
    elapsed_2 := time.Since(start_2)
    log.Printf("Decrypto took %s", elapsed_2)
    fmt.Println("plain text = ", c)

}
