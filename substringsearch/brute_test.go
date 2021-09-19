package substringsearch

import (
	"fmt"
	"testing"
)

// func TestBruteForce(t *testing.T) {
// 	st := "Hello I am inquiring about turtles"
// 	pat := "turt"
// 	fakePat := "turd"
// 	fmt.Println(BruteForce(pat,st))
// 	fmt.Println(BruteForce(fakePat,st))

// }

func TestCyclicRotation(t *testing.T) {
	s := "winterbreak"
	j := "breakwonter"
	fmt.Println("Cyclic Rot:",cycRot(s,j))
}	