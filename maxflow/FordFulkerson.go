package maxflow

import (
	"math"
)

type FordFolkerson struct {
	Marked map[int]bool
	EdgeTo map[int]FlowEdge
	Value float64
}

func (ff *FordFolkerson) InCut(v int) bool {
	return ff.Marked[v]
}

func RunFordFolkerson(g *FlowNetwork, s int, t int) FordFolkerson {
	ff := &FordFolkerson{
		Marked: make(map[int]bool),
		EdgeTo: make(map[int]FlowEdge),
		Value:0.0,
	}
	for ff.hasAugmentingPath(g,s,t) {
		bottle := math.Inf(1)
		for v := t; v != s; v = ff.EdgeTo[v].other(v) {
			bottle = math.Min(bottle,ff.EdgeTo[v].ResidualCapacityTo(v))
		}
		for v := t; v != s; v = ff.EdgeTo[v].other(v) {
			ed := ff.EdgeTo[v]
			ed.Flow = ff.EdgeTo[v].addResidualFlowTo(v,bottle)
		}
		ff.Value += bottle
	}
	return (*ff)
}


func (ff *FordFolkerson) hasAugmentingPath(g *FlowNetwork, s int, t int) bool {
	ff.Marked = make(map[int]bool)
	ff.EdgeTo = make(map[int]FlowEdge)

	queue := []int{}
	queue = append(queue, s)

	ff.Marked[s] = true
	for len(queue) != 0 {
		v := queue[0]
		queue = queue[1:]

		for _,e := range g.Adj[v] {
			w := e.other(v)
			res := e.ResidualCapacityTo(w)
			_,ok := ff.Marked[w]
			if res > 0 && !ok {
				ff.EdgeTo[w] = e
				ff.Marked[w] = true
				queue = append(queue,w)
			} 
		}
	}
	return ff.Marked[t]
}