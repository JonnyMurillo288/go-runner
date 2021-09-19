package trie

import "strconv"



func NewBinaryTrie(bin []string) *TrieSymbolTable {
	st := NewTrieSymbolTable(2)
	var t = 0
	for _,b := range bin {
		st.Put(b,t)
		t++
	}
	return st
}

func (t *TrieSymbolTable) SearchPre(pattern string, vis map[string]bool) (bool,map[string]bool) {
	return searchPre(t.Root,pattern,0,vis)
}

func searchPre(x *Node, pattern string, d int, visited map[string]bool) (bool,map[string]bool) {
	if x.Obj == nil {
		return true,visited
	}
	if d == len(pattern) {
		return true,visited
	}
	if visited[pattern[:d]] {
		return false,visited
	}
	visited[pattern[:d]] = true
	c := pattern[d]
	v,_ := strconv.ParseInt(string(c),10,64)
	return searchPre(x.Next[v],pattern,d+1,visited)
}