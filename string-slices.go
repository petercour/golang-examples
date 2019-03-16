// https://golangr.com/slices/
package main

import "fmt"

func main() {
    /* define a string */ 
    mystring := "Go programming"

    /* take slice */
    s := mystring[0:2]
    fmt.Println(s)
}