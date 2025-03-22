/*
Print => prints with default formatting
Println => adds space between args and newline at the end
*/

package main
import ("fmt")

func main() {
    i, j := "Hello", "World"
    fmt.Print(i, "\n")
    fmt.Print(j, " ")
    fmt.Println(i,j)
    fmt.Printf("%v\n",i)
    fmt.Printf("%T",j)
}
