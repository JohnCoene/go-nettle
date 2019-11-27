package nettle

import (
	"math"
)

func fa(x, k float64) float64 {
	return (x * x) / k
}

func fr(x, k float64) float64 {
	return (k * k) / x
}

func findNodes(edge Edge, nodes []NodeLayout) (nettle.NodeLayout, nettle.NodeLayout) {
	var found int = 0
	for _, n := range nodes {

		if n.id == edge.source {
			found = found + 1
			src := n
		}

		if n.id == edge.target {
			found = found + 1
			tgt := n
		}

		if found == 2 {
			break
		}

	}

	return src, tgt
}

func distance(x1, x2, y1, y2, float64) float64 {
	xDiff := x1 - x2
	yDiff := y1 - y2
	diff := float64(((xDiff * xDiff) + (yDiff * yDiff)))
	return math.Sqrt(diff)
}

// Layout Compute graph layout
func Layout(W, L, iter int, g GraphLayout) int {

	nNodes := float64(len(g.Nodes))
	area := float64((W * L))
	k := math.Sqrt(area / nNodes)

	for i := 0; i < iter; i++ {

		for _, u := range g.Nodes {
			u.disp = float64(0)

			for _, v := range g.Nodes {
				if u != v {
					d := distance(v.x, u.x, v.y, u.y)
					v.disp = v.disp + (d / math.Abs(d) * fr(math.Abs(d), k))
				}
			}

		}

		for _, e := range g.Edges {
			source, target := findNodes(e, g.Nodes)
			d := distance(source.x, target.x, source.y, target.y)
			source.disp = source.disp - (d/math.Abs(d))*fa(math.Abs(d), k)
			target.disp = target.disp - (d/math.Abs(d))*fa(math.Abs(d), k)
		}

	}

	return 0
}
