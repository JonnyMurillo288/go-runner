package substringsearch

import "fmt"


func BruteForce(pattern string, text string) int {
	for i := 0; i < len(text)-len(pattern); i++ {
		for j := 0; j < len(pattern);j++ {
			fmt.Printf("Comparing %v --> %v\n",string(text[j+i]),string(pattern[j]))
			if text[j+i] != pattern[j] {
				break
			}
			if j == len(pattern)-1 {
				return i
			}
		}
	}
	return len(text)
}