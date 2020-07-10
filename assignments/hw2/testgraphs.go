// Copyright ©2014 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hw2

import (
	"fmt"
	"math"

	"github.com/gonum/graph"
	"github.com/gonum/graph/simple"
)

func init() {
	for _, test := range ShortestPathTests {
		if len(test.WantPaths) != 1 && test.HasUniquePath {
			panic(fmt.Sprintf("%q: bad shortest path test: non-unique paths marked unique", test.Name))
		}
	}
}

// ShortestPathTests are graphs used to test the static shortest path routines in path: BellmanFord,
// DijkstraAllPaths, DijkstraFrom, FloydWarshall and Johnson, and the static degenerate case for the
// dynamic shortest path routine in path/dynamic: DStarLite.
var ShortestPathTests = []struct {
	Name              string
	Graph             func() graph.EdgeSetter
	Edges             []simple.Edge
	HasNegativeWeight bool
	HasNegativeCycle  bool

	Query         simple.Edge
	Weight        float64
	WantPaths     [][]int
	HasUniquePath bool

	NoPathFor simple.Edge
}{
	
	{
		Name:  "zero-weight |V|·cycle^(n/|V|) directed",
		Graph: func() graph.EdgeSetter { return simple.NewDirectedGraph(0, math.Inf(1)) },
		Edges: func() []simple.Edge {
			e := []simple.Edge{
				// Add a path from 0->4 of weight 4
				{F: simple.Node(0), T: simple.Node(1), W: 1},
				{F: simple.Node(1), T: simple.Node(2), W: 1},
				{F: simple.Node(2), T: simple.Node(3), W: 1},
				{F: simple.Node(3), T: simple.Node(4), W: 1},
			}
			next := len(e) + 1

			// Add n zero-weight cycles.
			const n = 100
			for i := 0; i < n; i++ {
				e = append(e,
					simple.Edge{F: simple.Node(next + i), T: simple.Node(i), W: 0},
					simple.Edge{F: simple.Node(i), T: simple.Node(next + i), W: 0},
				)
			}
			return e
		}(),

		Query:  simple.Edge{F: simple.Node(0), T: simple.Node(4)},
		Weight: 4,
		WantPaths: [][]int{
			{0, 1, 2, 3, 4},
		},
		HasUniquePath: false,

		NoPathFor: simple.Edge{F: simple.Node(4), T: simple.Node(5)},
	},
	{
		Name:  "zero-weight n·cycle directed",
		Graph: func() graph.EdgeSetter { return simple.NewDirectedGraph(0, math.Inf(1)) },
		Edges: func() []simple.Edge {
			e := []simple.Edge{
				// Add a path from 0->4 of weight 4
				{F: simple.Node(0), T: simple.Node(1), W: 1},
				{F: simple.Node(1), T: simple.Node(2), W: 1},
				{F: simple.Node(2), T: simple.Node(3), W: 1},
				{F: simple.Node(3), T: simple.Node(4), W: 1},
			}
			next := len(e) + 1

			// Add n zero-weight cycles.
			const n = 100
			for i := 0; i < n; i++ {
				e = append(e,
					simple.Edge{F: simple.Node(next + i), T: simple.Node(1), W: 0},
					simple.Edge{F: simple.Node(1), T: simple.Node(next + i), W: 0},
				)
			}
			return e
		}(),

		Query:  simple.Edge{F: simple.Node(0), T: simple.Node(4)},
		Weight: 4,
		WantPaths: [][]int{
			{0, 1, 2, 3, 4},
		},
		HasUniquePath: false,

		NoPathFor: simple.Edge{F: simple.Node(4), T: simple.Node(5)},
	},
	{
		Name:  "zero-weight bi-directional tree with single exit directed",
		Graph: func() graph.EdgeSetter { return simple.NewDirectedGraph(0, math.Inf(1)) },
		Edges: func() []simple.Edge {
			e := []simple.Edge{
				// Add a path from 0->4 of weight 4
				{F: simple.Node(0), T: simple.Node(1), W: 1},
				{F: simple.Node(1), T: simple.Node(2), W: 1},
				{F: simple.Node(2), T: simple.Node(3), W: 1},
				{F: simple.Node(3), T: simple.Node(4), W: 1},
			}

			// Make a bi-directional tree rooted at node 2 with
			// a single exit to node 4 and co-equal cost from
			// 2 to 4.
			const (
				depth     = 4
				branching = 4
			)

			next := len(e) + 1
			src := 2
			var i, last int
			for l := 0; l < depth; l++ {
				for i = 0; i < branching; i++ {
					last = next + i
					e = append(e, simple.Edge{F: simple.Node(src), T: simple.Node(last), W: 0})
					e = append(e, simple.Edge{F: simple.Node(last), T: simple.Node(src), W: 0})
				}
				src = next + 1
				next += branching
			}
			e = append(e, simple.Edge{F: simple.Node(last), T: simple.Node(4), W: 2})
			return e
		}(),

		Query:  simple.Edge{F: simple.Node(0), T: simple.Node(4)},
		Weight: 4,
		WantPaths: [][]int{
			{0, 1, 2, 3, 4},
			{0, 1, 2, 6, 10, 14, 20, 4},
		},
		HasUniquePath: false,

		NoPathFor: simple.Edge{F: simple.Node(4), T: simple.Node(5)},
	},
}
