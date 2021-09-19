package compression

import (
	"math"

	trie "github.com/Jonnymurillo288/go-runner/tries"
)

func LZWCompress(input string) []int {
	res := []int{}
	R := 256
	L := math.Pow(2,12)
	tst := trie.NewTST("",nil)
	for i := 0; i < R; i++ {
		tst.Put(""+string(rune(i)),i)
	}
	code := R + 1
	for len(input) > 0 {
		s := tst.LongestPrefixOf(input)
		get := tst.Get(s)
		res = append(res,get.(int))
		t := len(s)
		if t < len(input) && code < int(L) {
			code++
			tst.Put(input[0:t+1],code)
		}
		input = input[t:]
	}
	return res
}