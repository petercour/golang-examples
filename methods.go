// https://golangr.com/methods/
package main

import "fmt"

func main() {
   hello("go")
   hello("")
}


func hello(x1 string) {
   fmt.Printf( "Hello %s", x1);
}