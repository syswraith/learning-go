package main
import "fmt"

func main() {

	for i := range 100 {

		if i % 2 == 0 {
			fmt.Println("Fizz")
		} else if i % 3 == 0	{
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}

	}

}
