package main

import "fmt"

func main() {

	//decision dan looping
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Printf("fizz = [%v]\n", i)
		} else if i%5 == 0 {
			fmt.Printf("buzz = [%v]\n", i)
		} else {
			fmt.Println(i)
		}
	}
}
