package main;

import "fmt"

func main() {

	fmt.Println("normal for loop")
	for i := 0; i < 100; i++ { 
		fmt.Println(i) 
	}

	fmt.Println("python styled for loop")
	for i := range 10 {
		fmt.Println(i)
	}

	fmt.Println("for ;; loop")
	for {
		fmt.Print(0)
	}

}
