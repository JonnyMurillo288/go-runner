package substringsearch

import (
	"fmt"
	"testing"
)

func TestKMP(t *testing.T) {
	st := "Everything profound loves the mask"
	truePat := "prof"
	falsePat := "lives"
	fmt.Printf("Should return true: %v\n",KMP(st,truePat)!=len(st))
	fmt.Printf("Should return false: %v\n",KMP(st,falsePat)!=len(st))
}