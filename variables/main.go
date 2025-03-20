/*
variables are declared with the var keyword
you can only use the walrus operator WITHIN a function
*/

package main
import ("fmt")

var integer int = 32
var decimal float32 = 32.12
var str string = "Hello world"
var boolean bool = true

func main() {
    random_stuff := "wow"
    fmt.Printf("%d\n", integer)
    fmt.Printf("%s\n", random_stuff)
    fmt.Printf("%f\n", decimal)
    fmt.Printf("%s\n", str)
    fmt.Printf("%t\n", boolean)
}
