package graphs

import (
	"fmt"
	"log"
	"math"

	utils "github.com/Jonnymurillo288/GoUtils"
)

type Dijkstras struct {
	EdgeTo []Edge
	DistTo []float64
	PQ utils.IndexMinPQ
}

// 1. consider vertices in decreasing order from the s
// (non-tree vertex with lowest distTo Value)
// 2. add the vertex to the tree and relax all edges pointing from that vertex
func NewDijkstras(g *EdgeWeightDigraph, s int) Dijkstras {
	d := &Dijkstras{
		EdgeTo: make([]Edge,g.LV()+1),
		DistTo: make([]float64,g.LV()+1),
		PQ: utils.NewIndexMinPQ(g.LV()+1),
	}
	for v := 0; v < g.LV(); v++ {
		d.DistTo[v] = math.Inf(-1) // distance to v is infinity
	}
	d.DistTo[s] = 0.0

	d.PQ.Insert(s,0.0) // relax vertices in order of distance from s
	fmt.Println("aigoheow: ",g.Adj[0],g.Adj[249])
	for !d.PQ.IsEmpty() {
		v := d.PQ.DelMin() // first pass deletes the source
		fmt.Println("Looping through pq with min value of:",v)
		// second pass will delete the min from source to next vertice
		// loop through all vertices that are connected to the min
		for _,e := range g.Adj[v] {
			d.Relax(e)
		}
	}
	return (*d)
}


// relax the edge is the distance from v to w is shorter than distance to w
func (d *Dijkstras) Relax(e Edge) {
	v := e.From()
	w := e.To() 
	log.Printf("\n\nRelaxing From %v - %v",v,w)
	log.Println("EDGE:",e)
	log.Printf("\nCurrent distTo for %v : %v",w,d.DistTo[w])
	log.Printf("Potential DistTo: %v",d.DistTo[v]+e.Weight)
	if d.DistTo[w] > d.DistTo[v] + e.Weight {
		d.DistTo[w] = d.DistTo[v] + e.Weight
		d.EdgeTo[w] = e
		if d.PQ.Contains(w) {
			log.Printf("PQ contains %v\n",w)
			// if w in the Priority queue and distTo[w] less than current distTo[w]
			// decrese the key
			d.PQ.DecreaseKey(w, d.DistTo[w])
		} else {
			// otherwise insert the key into the PQ
			fmt.Printf("\nInserting %v into PQ",w)
			d.PQ.Insert(w, d.DistTo[w])
		}
	}
}