package lib

import "fmt"

func PrintHeader() {
	fmt.Printf("description | first element pointer | len | cap | slice\n")
}

func PrintSlice(desc string, s []int) {
	var ptr *int
	if len(s) > 0 {
		ptr = &s[0]
	}
	fmt.Printf("%11s | %21p | %3d | %3d | %v\n", desc, ptr, len(s), cap(s), s)
}
