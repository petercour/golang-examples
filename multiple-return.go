// https://golangr.com/multiple-return/
package main

import "fmt"

func values() (int, int) {
   return 2,4
}

func main() {
   x, y := values()
   fmt.Println(x)
   fmt.Println(y)
}