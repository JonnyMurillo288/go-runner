package trie

import "strconv"

type Node struct {
	Obj interface{}
	Next []*Node
}

type TrieSymbolTable struct {
	R int
	Root *Node
}

func newNode(obj interface{},r int) *Node {
	return &Node{
		Obj:obj,
		Next:make([]*Node,r), // each node holds list of all ascii values
	}
}

// create
func NewTrieSymbolTable(r int) *TrieSymbolTable {
	t := &TrieSymbolTable{
		R:r,
		Root: newNode(nil,r),
	}
	return t
}

func (t *TrieSymbolTable) Put(key string, val interface{}) {
	t.Root = t.put(t.Root,key,val,0)
}

func (t *TrieSymbolTable) put(x *Node, key string, value interface{}, d int) *Node {
	if x == nil {
		x = newNode(nil,t.R)
	}
	if d == len(key) {
		x.Obj = value
		return x
	}
	c := key[d]
	v,_ := strconv.ParseInt(string(c),10,64)
	x.Next[v] = t.put(x.Next[v],key,value,d+1)
	return x
}

func (t *TrieSymbolTable) contains(key string) bool {
	return t.Get(key) != nil 
}

func (t *TrieSymbolTable) Get(key string) interface{} {
	x := t.get(t.Root,key,0)
	if x.Obj == nil {
		return nil
	}
	return x.Obj
}

func (t *TrieSymbolTable) get(x *Node, key string, d int) *Node {
	if x == nil {
		return newNode(nil,t.R)
	}
	if d == len(key) {
		return x
	}
	c := key[d]
	return t.get(x.Next[c],key,d+1)
}