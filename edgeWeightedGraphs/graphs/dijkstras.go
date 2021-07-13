package graphs

import (
	"math"

	utils "github.com/Jonnymurillo288/GoUtils"
)

type Dijkstras struct {
	EdgeTo []*Edge
	DistTo []float64
	PQ utils.IndexMinPQ
}

// 1. consider vertices in decreasing order from the s
// (non-tree vertex with lowest distTo Value)
// 2. add the vertex to the tree and relax all edges pointing from that vertex
func NewDijkstras(g *EdgeWeightDigraph, s int) Dijkstras {
	d := &Dijkstras{
		EdgeTo: make([]*Edge,g.LV()+1),
		DistTo: make([]float64,g.LV()+1),
		PQ: utils.NewIndexMinPQ(g.LV()+1),
	}
	for v := 0; v < g.LV(); v++ {
		d.DistTo[v] = math.Inf(1) // distance to v is infinity
	}
	d.DistTo[s] = 0.0

	d.PQ.Insert(s,0.0) // relax vertices in order of distance from s
	for !d.PQ.IsEmpty() {
		v := d.PQ.DelMin()
		if v == -1 {
			return (*d)
		}
		for _,e := range g.Adj[v] {
			d.Relax(&e)
		}
	}
	return (*d)
}

func (d *Dijkstras) Relax(e *Edge) {
	v := e.From()
	w := e.To() 
	if d.DistTo[w] >= d.DistTo[v] + e.Weight {
		d.DistTo[w] = d.DistTo[v] + e.Weight
		d.EdgeTo[w] = e
		if d.PQ.Contains(w) {
			// if w in the Priority queue and distTo[w] less than current distTo[w]
			// decrese the key
			d.PQ.DecreaseKey(w, d.DistTo[w])
		} else {
			// otherwise insert the key into the PQ
			d.PQ.Insert(w, d.DistTo[w])
		}
	}
}