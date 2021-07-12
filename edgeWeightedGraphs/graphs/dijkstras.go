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

func NewDijkstras(g *EdgeWeightDigraph, s int) Dijkstras {
	d := &Dijkstras{
		EdgeTo: make([]*Edge,g.V()),
		DistTo: make([]float64,g.V()),
		PQ: utils.New(),
	}
	for v := 0; v < g.V(); v++ {
		d.DistTo[v] = math.Inf(1) // distance to v is infinity
	}
	d.DistTo[s] = 0.0

	d.PQ.Insert(s,0.0)
}

func (d *Dijkstras) Relax(e *Edge) {
	v := e.From()
	w := e.To() 
	if s.DistTo[w] > s.DistTo[v] + e.Weight {
		s.DistTo[w] = s.DistTo[v] + e.Weight
		s.EdgeTo[w] = e
		if pq.contains(w) {
			pq.DecreaseKey(w, d.DistTo[w])
		} else {
			pq.Insert(w, DistTo[w])
		}
	}
}