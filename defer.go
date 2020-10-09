package main
import "fmt"

func main() {
   defer FirstFunction()
    SecondFunction()


    fmt.Println("Defer will execute function in the last.")
}

func FirstFunction() {
    fmt.Println("This is FirstFunction")
}

func SecondFunction() {
    fmt.Println("This is SecondFunction")
}