// https://golangr.com/slices/
package main

import "fmt"

func main() {
    /* create a set/list of numbers */ 
    myset := []int{0,1,2,3,4,5,6,7,8}

    /* take slice */
    s := myset[0:4]
    fmt.Println(s)
}