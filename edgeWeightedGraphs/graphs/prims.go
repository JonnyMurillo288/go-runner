package graphs

var marked = make(map[int]bool)


// 1. Start with vertex 0
// 2. find the edge with minimum weight from vertex and as to MST
// 3. repeat until V - 1 
func (g *EdgeWeightGraph) PrimsLazy() []*Edge {
	var mst []*Edge
	var pq []*Edge
	for _,ed := range g.Edges {
		e := &ed
		pq = Add(pq, e)
	}
	for !isEmpty(pq) && len(mst) > g.LV()-1 {
		e := pq[0]
		pq = pq[1:]
		v := e.either()
		w := e.other(v)
		if marked[v] && marked[w] {
			continue
		}
		mst = append(mst,e)
		if !marked[v] {
			for _,p := range g.visit(v) {
				pq = Add(pq, p)
			}

		}
	}
	return mst
}

// 1. Maintain pq of verticies connected by edge to MST
// (Order this pq by weight of edge connecting v to MST)
// 2. Delete pq[0] (min weight) and add e = v-w to MST
// 3. Update pq by considering v-x
// 3.1 Ignore if x in MST
// decrease priority of x if v–x becomes shortest edge connecting x to MST
func (g *EdgeWeightGraph) PrimsEager() []*Edge {
	return []*Edge{}
}

// 
func (g *EdgeWeightGraph) visit(v int) []*Edge {
	var p []*Edge
	marked[v] = true
	for _,e := range g.Adj[v] {
		if !marked[e.other(v)] {
			p = append(p,&e)
		}
	}
	return p
}
