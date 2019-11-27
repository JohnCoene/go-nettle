package nettle

// Edge A set of edges
type Edge struct {
	source string
	target string
}

// NodeLayout Nettle node with position and displacement
type NodeLayout struct {
	id   string
	x    int
	y    int
	disp float64
}

// GraphLayout A graph
type GraphLayout struct {
	Nodes []NodeLayout
	Edges []Edge
}

// Graph A graph
type Graph struct {
	Nodes []string
	Edges []Edge
}

// CreateGraph Create a nettle graph
func (g Graph) CreateGraph() GraphLayout {

	nodes := make([]NodeLayout, 0)

	for _, n := range g.Nodes {
		nl := NodeLayout{
			id:   n,
			x:    0,
			y:    0,
			disp: 0,
		}
		nodes = append(nodes, nl)
	}

	gl := GraphLayout{
		Nodes: nodes,
		Edges: g.Edges,
	}

	return gl

}
