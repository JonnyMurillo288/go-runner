package graphs

type Edge struct {
	V,W int
	Weight float64
}

func (e *Edge) either() int {
	return e.V
}

func (e *Edge) other(v int) int {
	if v == e.V{
		return e.W
	} else {
		return e.V
	}
}

// -1: this edge weight is less
// 1: this edge weight is greater
// 0: both edge weights are the same
func (this *Edge) compareTo(that *Edge) int {
	if this.Weight < that.Weight {
		return -1
	} else if this.Weight > that.Weight {
		return 1
	} else {
		return 0
	}
}

// used to find directed edge
func(e *Edge) To() int {
	return e.W
}

func (e *Edge) From() int {
	return e.V
}

// ================================================================= //
// Edge Weighted Graph (not directed)
type EdgeWeightGraph struct {
	Edges []Edge
	Adj map[int][]Edge
}

func (e *EdgeWeightGraph) AddEdge(ed Edge) {
	var v,w int
	v = ed.either()
	w = ed.other(v)
	e.Adj[v] = append(e.Adj[v],ed)
	e.Adj[w] = append(e.Adj[w],ed)
}

func (e *EdgeWeightGraph) LV() int {
	var len int
	for _,_ = range e.Adj {
		len++
	}
	return len
}

func (e *EdgeWeightGraph) E() int {
	return len(e.Edges)
}


func NewWeightGraph() *EdgeWeightGraph {
	return &EdgeWeightGraph{
		Edges: make([]Edge,0),
		Adj: make(map[int][]Edge),
	}
}

// ================================================================= //
// Edge Weighted Digraph


type EdgeWeightDigraph struct {
	V int
	Adj map[int][]Edge
}

func NewWeightDigraph(v int) *EdgeWeightDigraph {
	return &EdgeWeightDigraph{
		V: v,
		Adj: make(map[int][]Edge),
	}
}

func (e *EdgeWeightDigraph) AddEdge(ed Edge) {
	v := ed.From()
	e.Adj[v] = append(e.Adj[v],ed)
}

func (e *EdgeWeightDigraph) LV() int {
	var len int
	for _,_ = range g.Adj {
		len++
	}
	return len
}


// ================================================================= //
//
// Work on this, create an MST from the algorithms and when to append them
// TODO: Create the Constructor to add to MST with the weights and edges
//
type MST struct {
	Edges []*Edge
	Weights []float64
}

func (m *MST) Add(e *Edge) {
	m.Edges = append(m.Edges,e)
}

func (m *MST) weights() []float64 {
	var w []float64
	for _,i := range m.Edges {
		w = append(w,i.Weight)
	}
	return w
}






// ================================================================= //

func inArr(arr []Edge, a Edge) bool {
	for _,p := range arr {
		if p == a {
			return true
		}
	}
	return false
}
