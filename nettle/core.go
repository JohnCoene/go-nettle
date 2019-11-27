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

func distance(minX, maxX, minY, maxY float64) float64 {
	xDiff := minX - maxX
	yDiff := minY - maxY
	diff := float64(((xDiff * xDiff) + (yDiff * yDiff)))
	return math.Sqrt(diff)
}

func findNodes(edge Edge, nodes []NodeLayout) (NodeLayout, NodeLayout, int, int) {
	var found, srcid, tgtid int
	var src, tgt NodeLayout
	for _, n := range nodes {

		if n.id == edge.Source {
			found = found + 1
			src = n
		}

		if n.id == edge.Target {
			found = found + 1
			tgt = n
		}

		if found == 2 {
			break
		}

	}

	return src, tgt, srcid, tgtid
}

// Layout Compute graph layout
func Layout(W, L, iter int, t float64, g GraphLayout) GraphLayout {

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
			source, target, srcid, tgtid := findNodes(e, g.Nodes)
			d := distance(source.x, target.x, source.y, target.y)
			source.disp = source.disp - (d/math.Abs(d))*fa(math.Abs(d), k)
			target.disp = target.disp - (d/math.Abs(d))*fa(math.Abs(d), k)
			g.Nodes[srcid].disp = source.disp
			g.Nodes[tgtid].disp = target.disp
		}

		for _, n := range g.Nodes {
			minDisp := float64(math.Min(n.disp, t))
			n.x = n.x + (n.disp/math.Abs(n.disp))*minDisp
			n.y = n.y + (n.disp/math.Abs(n.disp))*minDisp
			n.x = math.Min(float64(W/2), math.Max(float64(-W/2), n.x))
			n.y = math.Min(float64(L/2), math.Max(float64(-L/2), n.y))
		}

	}

	return g
}
