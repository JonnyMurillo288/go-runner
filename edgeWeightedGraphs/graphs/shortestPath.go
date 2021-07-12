package graphs

import (
	"fmt"

	"github.com/emirpasic/gods/stacks/arraystack"
)

type ShortestPath struct {
	Graph *EdgeWeightDigraph
	S int
	DistTo map[int]float64
	EdgeTo []*Edge
}

func FindShortestPath(g *EdgeWeightDigraph, s int) *ShortestPath {
	return &ShortestPath{
		Graph: g,
		S: s,
	}
}



func (s *ShortestPath) PathTo(v int) []*Edge {
	var res []*Edge
	path := arraystack.New()

	var e = s.EdgeTo[v]
	for {
		e = s.EdgeTo[e.From()]
		if e != nil {
			s.DistTo[e.From()] += e.Weight
			path.Push(e)
		} else {
			break
		}
	}

	it := path.Iterator()
	for it.Next() {
		r := it.Value()
		ed := &r // create this value as a pointer
		fmt.Println(ed)
	}
	return res
}

func (s *ShortestPath) Relax(e *Edge) {
	v := e.From()
	w := e.To()
	if s.DistTo[w] > s.DistTo[v] + e.Weight {
		s.DistTo[w] = s.DistTo[v] + e.Weight
		s.EdgeTo[w] = e
	}
} 

