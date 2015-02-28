package main

import (rsa "./rsa"
        "fmt"
       )

func main(){
    fmt.Println("Hello, playground")
    priv := rsa.GeneratePrivateKey()
    pub := rsa.GeneratePublicKey(priv)
    fmt.Println("n = ", pub)

}