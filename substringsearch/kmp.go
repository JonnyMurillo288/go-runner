package substringsearch

import "fmt"


func KMP(text string, pat string) int {
	R := 256
	N := len(text)
	M := len(pat)
	j := 1
	i := 0
	ind := 0
	dfa := make([][]int,R)
	for m := range dfa {
		dfa[m] = make([]int,M)
	}
	dfa[pat[0]][0] = 1
	for X := 0; j < M; j++ {
		for c := 0; c < R; c++ {
			dfa[c][j] = dfa[c][X]
		}
		dfa[pat[j]][j] = j+1
		X = dfa[pat[j]][X]
	}
	j = 0
	for i = 0; i < N && j < M; i++ {
		j = dfa[text[i]][j]
		ind = i
	}
	if j == M {
		return ind - M
	} else {
		return N
	}
}

func readDfa(dfa [][]int) ([]int,[]int) {
	var res,fin []int
	for i,d := range dfa {
		for j,r := range d {
			if r != 0 {
				res = append(res,i)
				fin = append(fin,j)
				fmt.Printf("dfa[%v][%v] = %v\n",i,j,r)
			}
		}
	}
	return res,fin
}