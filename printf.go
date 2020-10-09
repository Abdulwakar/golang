package main

import "fmt"

func main() {
    var name string = "Abdul Shaikh"
    const pi float64 = 3.14121231212394838
    win := true
    x :=5


    fmt.Println(len(name))
    fmt.Println(name + " is good person")
    fmt.Printf("%.3f \n", pi)
    fmt.Printf("%f \n", pi)
    fmt.Printf("%T \n", name)
    fmt.Printf("%t \n", win)
    fmt.Printf("%d \n", x)
    fmt.Printf("%b \b", 25)
    fmt.Printf("%c \n", 21)
    fmt.Printf("%e \n", pi)
}