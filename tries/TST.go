package trie

type TSTNode struct {
	Value interface{}
	C int 
	Left, Mid, Right *TSTNode
}

type TST struct {
	Root *TSTNode
}

func newTSTNode() *TSTNode{
	return &TSTNode{
		Value: nil,
	}
}

func NewTST(key string, val interface{}) *TST {
	t := &TST{
		Root:newTSTNode(),
	}
	t.Root = t.put(t.Root,key,val,0)
	return t
}

func (t *TST) put(x *TSTNode,key string, val interface{}, d int) *TSTNode {
	c := int(key[d])
	if x == nil {
		x = newTSTNode()
		x.C = c
	}
	if c < x.C {
		x.Left = t.put(x.Left,key,val,d)
	} else if c > x.C {
		x.Right = t.put(x.Right,key,val,d)
	} else if d < len(key)-1 {
		x.Mid = t.put(x.Mid,key,val,d+1)
	} else {
		x.Value = val
	}
	return x
}