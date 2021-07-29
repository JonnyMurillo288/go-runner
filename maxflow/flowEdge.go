package maxflow

type FlowEdge struct {
	V, W int
	Capacity float64
	Flow float64
}

func NewFlowEdge(v int, w int, cap float64) FlowEdge {
	return FlowEdge{
		V:v,
		W:w,
		Capacity: cap,
	}
}


func (f FlowEdge) other(vertex int) int {
	if vertex == f.V {
		return f.W
	} else if vertex == f.W {
		return f.V
	}
	return -1
}

func (f FlowEdge) ResidualCapacityTo(vertex int) float64{
	if vertex == f.V {
		return f.Flow
	} else if vertex == f.W {
		return f.Capacity - f.Flow
	}
	return -1.0
}

func (f FlowEdge) addResidualFlowTo(vertex int, delta float64) float64 {
	if vertex == f.V {
		f.Flow -= delta
	} else if vertex == f.W {
		f.Flow += delta
	}
	return f.Flow
}

func (f FlowEdge) To() int {
	return f.W
}

func (f FlowEdge) From() int {
	return f.V
}