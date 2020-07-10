package hw2

import (
	"github.com/gonum/graph"
	"fmt"
)
// Apply the bellman-ford algorihtm to Graph and return
// a shortest path tree.
//
// Note that this uses Shortest to make it easier for you,
// but you can use another struct if that makes more sense
// for the concurrency model you chose.
type distance struct{
	to_idx int
	dist_new float64
	from_idx int
	changed bool
}
func UpdateDist(chnl chan distance, u graph.Node, v graph.Node, path Shortest, g graph.Graph, w float64)	{
	//get weight from u to v
	from := path.indexOf[u.ID()]
	to := path.indexOf[v.ID()]
	//w := g.Edge(u,v).Weight()
	
	dist_updated := path.dist[from] + w

	if dist_updated < path.dist[to]{
		//path.set(to, dist_updated, from)
		chnl<- distance{to_idx: to, dist_new: dist_updated, from_idx: from, changed: true}
	}else{
		chnl<- distance{to_idx: 0, dist_new: 0, from_idx: 0, changed: false}
	}
		
}
func BellmanFord(s graph.Node, g graph.Graph) Shortest {
	if !g.Has(s) {
		return Shortest {from: s}
	}
	var weight Weighting
	if wg, ok := g.(graph.Weighter); ok {
		weight = wg.Weight
	} else {
		weight = UniformCost(g)
	}

	nodes := g.Nodes()
	path := newShortestFrom(s, nodes)
	path.dist[path.indexOf[s.ID()]] = 0
	//fmt.Println(g.Nodes())

changed := true

for i:= 0; (i < (len(nodes)-1))&&(changed!=false); i++{
	changed = false
	for _,u := range g.Nodes() {
		chnl := make(chan distance)
		fromnodes:= g.From(u)
		for _, v := range g.From(u){
			w, ok := weight(u, v)
			if !ok{
				panic("Panic: incorrect weight")
			}
			go UpdateDist(chnl, u, v, path, g, w)
		}
		for j:=0; j < len(fromnodes); j++{
			newpath := <- chnl
				changed = newpath.changed
				if(changed == true){
					if newpath.dist_new < 0{
						panic("PANIC! at the disco. Negative Weight")
						//fmt.Println(v)
					}
					path.set(newpath.to_idx, newpath.dist_new, newpath.from_idx)
				}

		}
	}
	
}//END for all nodes 
fmt.Println("BELLMAN FORD: ", path.dist)
	return path
}

// Apply the delta-stepping algorihtm to Graph and return
// a shortest path tree.
//
// Note that this uses Shortest to make it easier for you,
// but you can use another struct if that makes more sense
// for the concurrency model you chose.


