package graphs

type Edge struct {
	v,w int
	weight float64
}

func (e *Edge) either() int {
	return e.v
}

func (e *Edge) other(v int) int {
	if v == e.v{
		return e.w
	} else {
		return e.v
	}
}

// -1: this edge weight is less
// 1: this edge weight is greater
// 0: both edge weights are the same
func (this *Edge) compareTo(that *Edge) int {
	if this.weight < that.weight {
		return -1
	} else if this.weight > that.weight {
		return 1
	} else {
		return 0
	}
}

// ================================================================= //

type EdgeWeightGraph struct {
	Edges []Edge
	Adj map[int][]Edge
}

func (e *EdgeWeightGraph) addEdge(ed Edge) {
	var v,w int
	v = ed.either()
	w = ed.other(v)
	e.Adj[v] = append(e.Adj[v],ed)
	e.Adj[w] = append(e.Adj[w],ed)
}

func (e *EdgeWeightGraph) V() int {
	var len int
	for _,_ = range e.Adj {
		len++
	}
	return len
}

func (e *EdgeWeightGraph) E() int {
	return len(e.Edges)
}

func (e *EdgeWeightGraph) edges() []Edge {
	var edges []Edge 
	for _,edge := range e.Adj {
		for _,ed := range edge {
			if !inArr(edges,ed) {
				edges = append(edges,ed)
			}
		}
	}
	return edges
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

func (m *MST) add(e *Edge) {
	m.Edges = append(m.Edges,e)
}

func (m *MST) edges() []*Edge {
	return m.Edges
}

func (m *MST) weights() []float64 {
	var w []float64
	for _,i := range m.Edges {
		w = append(w,i.weight)
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