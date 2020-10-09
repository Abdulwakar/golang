package main
import "fmt"

func main() {
    var EvenNum[5]  int
    EvenNum[0] = 0
    EvenNum[1] = 10
    EvenNum[2] = 20
    EvenNum[3] = 20
    EvenNum[4] = 30

    fmt.Println(EvenNum[4])

    fmt.Println("Below is an another way to define array")

    myNums := [5]int{1,2,4,5,6}
    fmt.Println(myNums[2])
    fmt.Println("Using of for loop")
    for i, value := range myNums {
        fmt.Println(value, i)
    }

    fmt.Println("Slice array")
    numSlice := []int{5,4,3,2,1}
    sliced := numSlice[2:5]
    fmt.Println(sliced)

    slice2 := make([]int,2)
    copy(slice2, numSlice)
    fmt.Println(slice2)
    fmt.Println(numSlice)

    slice3 := append(numSlice, 10,20)
    fmt.Println(slice3)

}