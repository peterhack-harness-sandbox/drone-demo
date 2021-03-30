package main

import "fmt"
//import "time"

func Hello() string {
    //return
    return "Hello, world"
    //change comment 1
    //time.Sleep(5 * time.Minute)
    
}

func main() {
    
    for {
    fmt.Println(Hello())
    }
}
