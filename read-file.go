// https://golangr.com/read-file/
package main

import (
    "fmt"
        "io/ioutil"
	)

func main() {
    b, err := ioutil.ReadFile("read.go")
    // can file be opened?
    if err != nil {
      fmt.Print(err)
    }

    // convert bytes to string
    str := string(b)

    // show file data
    fmt.Println(str)    
}