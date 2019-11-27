package nettle

import "math/rand"

// Edge A set of edges
type Edge struct {
	Source string
	Target string
}

// NodeLayout Nettle node with position and displacement
type NodeLayout struct {
	id   string
	x    float64
	y    float64
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
			x:    rand.Float64(),
			y:    rand.Float64(),
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
