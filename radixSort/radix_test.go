package radixsort

import (
	"fmt"
	"testing"
)

type v int

func TestKeyIndex(t *testing.T) {
	var i []interface{}
	a := []int{2345,325,251,46,5,6,547,54,7,2345,164,6,57,54,23,5,324652,56}
	for _,b := range a {
		i = append(i, b)
	}
	ki := newKeyIndex(i)
	fmt.Println(ki.sort())
}

// func TestJava(t *testing.T) {
// 	for i := 0; i < 5; i++ {
// 		java(i)
// 	}
// }