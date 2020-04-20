package main

import "fmt"

func printSlice(s string, x []int) {
	fmt.Printf("Slice %s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func main() {
	fmt.Println("Go Slice Example: Copy-on-Capacity-Overload\n")

	//create a slice of capacity 7
	a := make([]int, 7)
	b := a[2:5]
	printSlice("a", a)
	fmt.Println("Slice b refers to the 2, 3, 4 indices in slice a. Hence, the capacity is 5 (= 7-2).")
	fmt.Println("b := a[2:5]")
	printSlice("b", b)
	

	fmt.Println("\nModifying slice b, also modifies a, since they are pointing to the same underlying array.")
	fmt.Println("b[0] = 9")
	b[0] = 9
	printSlice("a", a)
	printSlice("b", b)
	

	fmt.Println("\nAppending 1 to slice b. Overwrites a.")
	b = append(b, 1)
	printSlice("a", a)
	printSlice("b", b)
	

	fmt.Println("\nAppending 2 to slice b. Overwrites a.")
	b = append(b, 2)
	printSlice("a", a)
	printSlice("b", b)
	

	fmt.Println("\nAppending 3 to slice b. Here, a new copy is made as the capacity is overloaded.")
	b = append(b, 3)
	printSlice("a", a)
	printSlice("b", b)
	

	fmt.Println("\nVerifying slices a and b point to different underlying arrays after the capacity-overload in the previous step.")
	fmt.Println("b[1] = 8")
	b[1] = 8
	printSlice("a", a)
	printSlice("b", b)
	return
}