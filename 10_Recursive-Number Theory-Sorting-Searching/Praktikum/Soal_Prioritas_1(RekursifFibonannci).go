package main

import "fmt"

func fibonanci (number int)int {
    if number <= 1{
        return number
    }
    return fibonanci(number-1) + fibonanci(number-2)
}
func main(){
    fmt.Println(fibonanci(0))
    fmt.Println(fibonanci(2))
    fmt.Println(fibonanci(9))
    fmt.Println(fibonanci(10))
    fmt.Println(fibonanci(12))
}