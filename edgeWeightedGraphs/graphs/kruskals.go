package graphs

import (
	utils "github.com/Jonnymurillo288/GoUtils"
)

// utils.Graph is a separate graph we can use to find cycles
func createUtilsGraph(m []*Edge) *utils.Graph {
	var g = &utils.Graph{
		Edges: make(map[int][]int),
	}
	for _,e := range m {
		v := e.either()
		w := e.other(v)
		g.AddEdge(v,w)		
	}
	return g
}

// add the edge to the priority queue and sort by weight
func add(arr []*Edge, a *Edge) []*Edge {
	var weightMap = make(map[float64]*Edge)
	var weights []float64
	var res []*Edge // results ordered list of Edges sorted by weight
	arr = append(arr,a) // add edge to the arr

	for _,f := range arr{ // loop through the arr
		weights = append(weights,f.weight) // add each weight to list of weights
		weightMap[f.weight] = f // weightMap[weight] = edge
	}
	weights = utils.SortFloat(weights) // sort the weights from least to greatest
	for _,w := range weights { 
		e := weightMap[w]
		res = append(res,e)
	}
	return res
}

func isEmpty(pq []*Edge) bool {
	var len int
	for _,_ = range pq {
		len++
	}
	if len == 0 {
		return true
	}
	return false
}


// 1. Sort edges by weight
// 2. Add the edge to MST UNLESS it creates a cycle
// 3. Use union find to see if we will create a cycle
func (g *EdgeWeightGraph) Kruskal() []*Edge {
	var mst []*Edge
	var pq []*Edge
	for _,ed := range g.Edges {
		e := &ed
		pq = add(pq, e)
	}
	uf := createUtilsGraph(mst)
	for !isEmpty(pq) && len(mst) < uf.V()-1 {
		e := pq[0]
		pq = pq[1:]
		v := e.either()
		w := e.other(v)
		gr := uf
		gr.Union(v,w)
		if !gr.IsCyclic() {
			uf.Union(v,w)
			mst = append(mst,e)
		}
	}
	return mst
}