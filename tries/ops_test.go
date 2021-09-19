package trie

import (
	"fmt"
	"testing"
)

func TestMatch(t *testing.T) {

	binary := []string{"01","10","0010","1111"} // no prefix
	binary2 := []string{"01","10","0010","10100"}  // prefix
	var f,tr bool
	bt := NewBinaryTrie(binary)
	bt2 := NewBinaryTrie(binary2)
	visited := make(map[string]bool)
	for _,p := range binary {
		f,visited = bt.SearchPre(p,visited)
	}
	visited = make(map[string]bool)
	for _,p := range binary2 {
		tr,visited = bt2.SearchPre(p,visited)
	}
	fmt.Println(f,tr)
}