package main
import "fmt"

func main() {
       x, y := 5, 6
       fmt.Println(add(x, y))
       fmt.Println(sub(x, y))
       fmt.Println("X + Y: ", x+y)

       fmt.Println("Example for recursion")

       num := 4
       fmt.Println(recursion(num))

       fmt.Println("Get value from Array and sum it")
       fmt.Println(ArrayAdd(10,20,30,40,50, 60))


}

func add(num1, num2 int) int{
    return num1+num2
}

func sub(num1, num2 int) int{
    return num2-num1
}

func recursion(num int) int {
    if num == 0 {
        return 1
    }
    return num * recursion(num - 1)
}

func ArrayAdd(args ...int) int{
    sum := 0
    for _, value := range args{
        sum += value
    }
    return sum
}