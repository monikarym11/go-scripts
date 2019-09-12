package main

import (
	"sort"
	"fmt"
	"math"	
)
func makeArrayConsecutive2(statues []int) int {
    sort.Ints(statues)
    n := 0
    for i, element := range statues[:len(statues)-1] {
        n = n + (statues[i+1]-element-1)
    }
    return n
}
func arrayReplace(inputArray []int, elemToReplace int, substitutionElem int) []int {

    for i, el := range inputArray {
        if el == elemToReplace {
            inputArray[i] = substitutionElem 
        }
    }
    return inputArray
}
func arrayMaximalAdjacentDifference(inputArray []int) int {
    max := 0
    for i := 1; i < len(inputArray); i++ {
        if math.Abs(float64(inputArray[i] - inputArray[i-1])) > float64(max) {
            max = int(math.Abs(float64(inputArray[i] - inputArray[i-1])))
        }        
    }    
    return max    
}
func arrayChange(inputArray []int) int {
    n := 0
    for i := 1; i < len(inputArray); i++ {
        for inputArray[i] <= inputArray[i-1] {
            inputArray[i]++
            n++
        }
    }
    return n
}
func main(){
	statues := []int{6, 2, 3, 8}
	fmt.Printf("Function makeArrayConsecutive2([6, 2, 3, 8]) - expected: 3, got: %v",makeArrayConsecutive2(statues))
	inputArray := []int{1, 2, 1}
	elemToReplace := 1
	substitutionElem := 3
	fmt.Printf("\nFunction arrayReplace([1, 2, 1], 1, 3) - expected: [3 2 3], got: %v", arrayReplace(inputArray, elemToReplace, substitutionElem))
	inputArray = []int{2, 4, 1, 0}
	fmt.Printf("\nFunction arrayMaximalAdjacentDifference([2, 4, 1, 0]) - expected: 3, got: %v", arrayMaximalAdjacentDifference(inputArray))
	inputArray = []int{1, 1, 1}
	fmt.Printf("\nFunction arrayChange([1, 1, 1]) - expected: 3, got: %v", arrayChange(inputArray))
}
