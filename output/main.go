/*
Printf => prints with default formatting
Println => adds space between args and newline at the end
Print => prints each args as is
*/

package main
import ("fmt")

func main() {
    i, j := "Hello", "World"
    fmt.Print(i, "\n")
    fmt.Print(j, " ")
    fmt.Println(i,j)
    fmt.Printf("%v %v\n", i, 1000) // prints the default value
    fmt.Printf("%T",j) // prints the type
}
