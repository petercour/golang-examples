// https://golangr.com/maps/
package main

import "fmt"

func main() {
    elements := make(map[string]string)
    elements["O"] = "Oxygen"
    elements["Ca"] = "Calcium"
    elements["C"] = "Carbon"

    fmt.Println(elements["C"])
}