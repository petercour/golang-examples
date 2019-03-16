// https://golangr.com/write-file/
// Write file in go. You don't need to set any flags.
// This will overwrite the file if it already exists
package main

import "os"

func main() {
    file, err := os.Create("file.txt")

    if err != nil {
        return
    }
    defer file.Close()

    file.WriteString("write file in golang")
}