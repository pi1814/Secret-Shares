package main

import (
    "fmt"
    "github.com/codahale/sss"
    "os"
    "strconv"
    "encoding/hex"

)

func main() {


    input_msg:= "AWESOME" 

               
    secretshares1:=6
    noofshares1:=3

    args := len(os.Args[1:])

    if (args>0) {input_msg= string(os.Args[1])}
        if (args>1) {secretshares1,_ = strconv.Atoi(os.Args[2])}
        if (args>2) {noofshares1,_ = strconv.Atoi(os.Args[3])}

    n := byte(secretshares1)                
    k := byte(noofshares1)   

    fmt.Printf("Input Message:\t%s\n\n",input_msg)
    fmt.Printf("Policy. Any %d from %d\n\n",noofshares1,secretshares1)

    if (noofshares1>secretshares1) { 
        fmt.Printf("Cannot do this, as k greater than n")
        os.Exit(0)
    }

    shares, _:= sss.Split(n, k, []byte(input_msg)) 


    subset := make(map[byte][]byte, k)
    for x, y := range shares { 
        fmt.Printf("Share:\t%d\t%s\n",x,hex.EncodeToString(y))
        subset[x] = y
        if len(subset) == int(k) {
            break
        }
    }

    reconstructed := string(sss.Combine(subset))
    fmt.Printf("\nReconstructed: %s\n",reconstructed)

}

