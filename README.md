# go-nettle

[![Go Report Card](https://goreportcard.com/badge/github.com/JohnCoene/go-nettle)](https://goreportcard.com/report/github.com/JohnCoene/go-nettle) 

A foray into graph layout algorithms with golang. Implements the Barycentric Method described [here](http://cs.brown.edu/people/rtamassi/gdhandbook/chapters/force-directed.pdf). 

## Install

```bash
go get github.com/go-nettle/nettle
```

## Example

```go
package main

import (
	"fmt"
	"go-nettle/nettle"
)

func main() {

	var nodes = []string{"India", "Canada", "Japan"}

	edges := make([]nettle.Edge, 3)
	edges[0] = nettle.Edge{
		Source: "India",
		Target: "Canada",
	}
	edges[1] = nettle.Edge{
		Source: "Canada",
		Target: "Japan",
	}
	edges[2] = nettle.Edge{
		Source: "Japan",
		Target: "India",
	}

	g := nettle.Graph{
		Nodes: nodes,
		Edges: edges,
	}

	gl := g.CreateGraph()

	gl = nettle.Layout(100, 100, 1, 5, gl)

	fmt.Println(gl)
}
```
