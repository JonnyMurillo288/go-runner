package trie

import "fmt"

// Extention operations for the TrieSymbolTable

// return an iterable of the keys
func (t *TrieSymbolTable) keys() []string {
	q := []string{}
	return collect(t.Root,"",q)
}

// return iterable of keys with prefix of PARAM: s
func (t *TrieSymbolTable) keysWithPrefix(s string) []string {
	q := []string{}
	x := t.get(t.Root,s,0) //move down the tree to start at out prefix point
	return collect(x,s,q)
}

// return iterable of keys that match PARAM: s
// PARAM: s - `.` is a wildcard
func (t *TrieSymbolTable) keysThatMatch(s string) []string {
	return findMatch(t.Root,s,[]string{},0)
}

func (t *TrieSymbolTable) longestPrefixOf(s string) string {
	length := search(t.Root,s,0,0)
	return s[0:length]
}

func search(x *Node, query string, d int, length int) int {
	if x.Obj == nil {
		return length
	}
	if x.Obj != nil {
		length = d
	}
	if d == len(query) {
		return length
	}
	c := query[d]
	return search(x.Next[c],query,d+1,length)
}

// traverse the TrieSymbolTable for keys
func collect(x *Node, prefix string, queue []string) []string {
	if x.Obj == nil {
		return queue
	}
	if x.Obj != nil {
		queue = append(queue, prefix)
	}
	for c := 0; c < 256; c++ {
		queue = collect(x.Next[c],prefix+fmt.Sprint(c),queue)
	}
	return queue
}

// search the trie and if the the next node matches match[d] then continue searching
// if node == nil then we reached the end of the word and can add to matched
func findMatch(x *Node, match string, matched []string, d int) []string {
	if d > len(match) {
		return matched
	}
	c := match[d]
	if x.Obj == nil {
		fmt.Println(x.Next)
	}
	if c == '.' {
		for _,node := range x.Next {
			if node.Obj != nil {
				matched = append(matched,node.Obj.(string))
			}
		}
	} else if string(c) == x.Obj.(string) {
		matched = append(matched,x.Obj.(string))
		return findMatch(x.Next[c],match,matched,d+1)
	}
	matched = findMatch(x.Next[c],match,matched,d)

	return matched
}