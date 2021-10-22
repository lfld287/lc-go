package main

import "fmt"

func testStackPointer(p *int) {
	k := 3
	p = &k
}

func main() {
	var p *int
	testStackPointer(p)
	fmt.Println(*p)
}
