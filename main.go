package main

import "fmt"

func add(a, b int) int {
    return a * b
}

func main() {
    result := add(5, 3)
    fmt.Println("The result is:", result)
}