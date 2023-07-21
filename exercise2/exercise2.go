package main

import (
    "fmt"
    "github.com/diegovanne/go23/exercise2/services/sorting"
)

func main() {
    intArr := []int{5, 2, 8, 1, 9, 3}
    sorting.SortInts(intArr)
    fmt.Println("Sorted integer array:", intArr)

    floatArr := []float64{3.14, 1.1, 0.5, 2.71, 1.618}
    sorting.SortFloats(floatArr)
    fmt.Println("Sorted float64 array:", floatArr)

    stringArr := []string{"apple", "banana", "orange", "grape", "pear"}
    sorting.SortStrings(stringArr)
    fmt.Println("Sorted string array:", stringArr)
}
