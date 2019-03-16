// https://golangr.com/random-numbers/
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func random(min int, max int) int {
    return rand.Intn(max-min) + min
}

func main() {
    rand.Seed(time.Now().UnixNano())
    randomNum := random(0, 10)
    fmt.Printf("Random number: %d\n", randomNum)
}