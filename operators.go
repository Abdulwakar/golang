package main
import "fmt"

func main() {
    x, y := 15, 5

    fmt.Println("X + Y =", x+y)
    fmt.Println("X - Y =", x-y)
    fmt.Println("X * Y =", x*y)
    fmt.Println("X / Y =", x/y)
    fmt.Println("X % Y =", x%y)

    isbool := true
    hate := false

    fmt.Println(isbool && hate)
    fmt.Println(isbool || hate)
    fmt.Println(!isbool)
}