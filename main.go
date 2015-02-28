package main

import (rsa "./rsa"
        "fmt"
       )

func main(){
    fmt.Println("Hello, playground")
    priv , pub := rsa.GenerateKeys()

    fmt.Println("priv = ", priv)
    fmt.Println("pub = ", pub)

}