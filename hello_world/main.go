/* 
- In Go every program is part of a package. So the program below belongs to the "main" package (its like naming the packages)
- "{" cannot be on a new line
- fmt = format
*/

package main
import ("fmt")

func main() {
    fmt.Println("Hello world!") // this doesnt allow format specifiers
    fmt.Print("Hello world!") // nor does this
    fmt.Printf("Hello world!") // only this
}
