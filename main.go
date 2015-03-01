package main

import (rsa "./rsa"
        "fmt"
       )

func main(){
    fmt.Println("Hello, playground")
    priv , pub := rsa.GenerateKeys()

    fmt.Println("priv = ", priv)
    fmt.Println("pub = ", pub)
    sample := []byte("myString")
    ciphered := rsa.Cipher(pub,sample)
    ciphered_bytes := []byte(ciphered)
    result := rsa.Decipher(priv, ciphered_bytes)
    fmt.Println("result = ", result)

}