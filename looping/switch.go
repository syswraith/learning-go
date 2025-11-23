package main
import "fmt"

func main() {
	for i := range 5 {
		switch i {
		case 1, 3:
			fmt.Println("one")
		case 2, 4:
			fmt.Println("two")
		default:
			fmt.Println("number")
		}
	}
}
