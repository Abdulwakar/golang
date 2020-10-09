package main

import "fmt"

func main() {
 age := 23

 if age > 18 {
    fmt.Println("Yes, you can vote!")
 } else {
    fmt.Println("No, You can't vote")
 }

fmt.Println ("Below is an exmple for switch case")
 switch age {
   case 18: fmt.Println ("You can apply for VoterID: ")
   case 19: fmt.Println ("you have VoterID:")
   default: fmt.Println ("confirm you age first")
 }


}