package maxflow

type FlowNetwork struct {
	V int
	Adj map[int][]FlowEdge
}

func NewFlowNetwork(v int) *FlowNetwork {
	return &FlowNetwork{
		V:v,
		Adj: make(map[int][]FlowEdge),
	}
}

func (f *FlowNetwork) AddEdge(e FlowEdge) {
	v := e.From()
	w := e.To()
	f.Adj[v] = append(f.Adj[v],e)
	f.Adj[w] = append(f.Adj[w],e)
}
