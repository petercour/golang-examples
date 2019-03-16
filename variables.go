// https://golangr.com/variables/
package main

import "fmt"

func main() {
   apple := 3.0
   bread := 2.0
   price := apple + bread

   fmt.Printf("")
   fmt.Printf("Price: 	 %f",price)
   vat := price * 0.15
   fmt.Printf("VAT: 	 %f",vat)
   total := vat + price
   fmt.Printf("Total: 	 %f",total)
   fmt.Printf("")
}