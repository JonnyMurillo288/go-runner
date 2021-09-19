package substringsearch

import "fmt"


func cycRot(s string, t string) bool {
	v := string(s[0])
	var set,j int
	for i,c := range t {
		if string(c) == v {
			set = i
		}
	}
	for i := set; i < len(t)-1 && j<len(s)-1; i++ {
		fmt.Println(i,j)
		fmt.Println(len(t))
		if t[i] != s[j] {
			fmt.Printf("Comparing %v --> %v\n",string(t[i]),string(s[j]))
			return false
		}
		j++
	}
	return true
}