package cmd

import (
    "flag"
	"fmt"
	"log"
	"strconv"
	"strings"
	"github.com/diegovanne/go23/exercise2/internal/sorter"
)

func Execute() {
	intFlag := flag.Bool("int", false, "Sort integer array")
	floatFlag := flag.Bool("float", false, "Sort float array")
	stringFlag := flag.Bool("string", false, "Sort string array")
	mixFlag := flag.Bool("mix", false, "Sort mixed array")
	flag.Parse()

	if flag.NFlag() != 1 {
		fmt.Println("Usage: go run main.go [-int | -float | -string | -mix] <elements>")
		fmt.Println("Example 1: go run main.go -int 5 2 10 1")
		fmt.Println("Example 2: go run main.go -string apple orange banana")
		fmt.Println("Example 3: go run main.go -mix 5.5 apple 2.7 orange 3 banana")
		return
	}

	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("Error: No input provided.")
	}

	switch {
	case *intFlag:
		intArr, err := ParseIntArray(args)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		sorter.SortInts(intArr)
		printIntArray(intArr)

	case *floatFlag:
		floatArr, err := ParseFloatArray(args)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		sorter.SortFloats(floatArr)
		printFloatArray(floatArr)

	case *stringFlag:
		sorter.SortStrings(args)
		printStringArray(args)

	case *mixFlag:
		mixed, err := ParseMixedArray(args)
		if err != nil {
			fmt.Println("Error while parsing mixed values:", err)
			return
		}

		sorter.SortMixedArray(mixed)
		printMixedArray(mixed)

	default:
		log.Fatal("Error: Invalid flag provided.")
	}
}

func ParseIntArray(args []string) ([]int, error) {
	var intArr []int
	for _, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			return nil, err
		}
		intArr = append(intArr, num)
	}
	return intArr, nil
}

func ParseFloatArray(args []string) ([]float64, error) {
	var floatArr []float64
	for _, arg := range args {
		num, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			return nil, err
		}
		floatArr = append(floatArr, num)
	}
	return floatArr, nil
}

func ParseMixedArray(args []string) ([]interface{}, error) {
	result := make([]interface{}, len(args))
	for i, arg := range args {
		if val, err := strconv.Atoi(arg); err == nil {
			result[i] = val
		} else if val, err := strconv.ParseFloat(arg, 64); err == nil {
			result[i] = val
		} else {
			result[i] = arg
		}
	}
	return result, nil
}

func printIntArray(arr []int) {
	fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(arr)), " "), "[]"))
}

func printFloatArray(arr []float64) {
	fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(arr)), " "), "[]"))
}

func printStringArray(arr []string) {
	fmt.Println(strings.Join(arr, " "))
}

func printMixedArray(arr []interface{}) {
	var elements []string
	for _, v := range arr {
		elements = append(elements, fmt.Sprintf("%v", v))
	}
	fmt.Println(strings.Join(elements, " "))
}