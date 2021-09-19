package compression

import (
	"fmt"
	"testing"
)


func TestPrefix(t *testing.T) {
	length := 0
	q := "I like to eat ass. Devin like to eat ass. ass tastes nice and i like it"
	og := len(q)*8
	res := LZWCompress(q)
	for _,num := range res {
		st := fmt.Sprint(num)
		length += len(st)
	}
	fmt.Println("Original Length:",og)
	fmt.Println("Compressed Length:",length)
	fmt.Println("Ratio:",float64(length/og))
}