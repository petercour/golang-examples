// https://golangr.com/variadic-functions/
package main

import "fmt"

func sum(numbers ...int) {
    total := 0
    for _, num := range numbers {
        total += num
    }
    fmt.Println(total)
}

func main() {
    sum(2,3,4)
    sum(1,1)
    sum(2,3,4)
    sum(1,2,3,4,5,6,7)
}