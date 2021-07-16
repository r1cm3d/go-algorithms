package main

import (
	"errors"
	"fmt"
	"log"
)

//   1    carry
// 1 2 2
// 1 3 9
// 2 6 1

// 1 1    carry
// 1 9 9
// 1 3 9
// 3 3 8
func main() {
	a := [3]int{1, 2, 2}
	b := [3]int{1, 3, 9}
	r, err := add(a[:], b[:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(r)
	a = [3]int{1, 9, 9}
	b = [3]int{1, 3, 9}
	r, err = add(a[:], b[:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
}

func add(a, b []int) ([]int, error) {
	const base = 10
	var (
		carry int
		r []int
	)

	length := len(a)
	if length != len(b) {
		return r, errors.New("a and b must have the same length")
	}

	r = make([]int, 3)
	for i := length-1; i >= 0; i-- {
		r[i] = a[i] + b[i] + carry

		carry = 0
		if r[i] >= base {
			r[i] = r[i] - base
			carry = 1
		}
	}

	return r, nil
}
