// https://golangr.com/struct/
package main

import "fmt"

type Person struct {
   name string
   job string
}
func main() {
   var aperson Person
    
   aperson.name = "Albert"
   aperson.job = "Professor"
   
   fmt.Printf( "aperson.name =  %s\n", aperson.name)
   fmt.Printf( "aperson.job  =  %s\n", aperson.job)
}