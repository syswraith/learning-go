package main
import "fmt"

func main() {

	// var varname size type
	var arr1 [5] int
	arr1[1] = 10
	for i := range len(arr1) {
		fmt.Println(arr1[i])
	}
	fmt.Println(arr1)

	/// ... infers the length of the array
	arr2 := [...] int {1,2,3,4,5}
	fmt.Println(arr2)

	
}
