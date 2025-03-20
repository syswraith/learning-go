/*
constants are typed and untyped, and declared with the const keyword
in untyped, the type is inferred from the declaration
*/

package main
import ("fmt")

const integer int = 32
const decimal = 32.12

func main() {
    fmt.Printf("%d\n", integer)
    fmt.Printf("%f\n", decimal)
}
