package main

import "fmt"

func main() {
    var num1, num2 float64
    num1=2
    num2=4
    sum := num1 + num2
    difference := num1 - num2
    product := num1 * num2
    quotient := num1 / num2
    fmt.Printf("\nSum: %.2f\n", sum)
    fmt.Printf("Difference: %.2f\n", difference)
    fmt.Printf("Product: %.2f\n", product)
    fmt.Printf("Quotient: %.2f\n", quotient)
}
