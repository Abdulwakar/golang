package main

import "fmt"

func main() {
 for i:= 1; i<=3; i++ {
    fmt.Println(i)
 }

 fmt.Println("Another way to define loops")

 j:= 10
    for j>=8 {
       fmt.Println(j)
       j--
    }

 fmt.Println("nested for loops")

 for a:=1; a<10; a++ {
    for b:=1; b<a; b++ {
        fmt.Printf("*")
    }
    fmt.Println()
 }
}