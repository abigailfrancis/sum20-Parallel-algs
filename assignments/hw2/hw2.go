package hw2

import (
	"github.com/gonum/graph"
	"fmt"
	"sync"
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
func UpdateDist(chnl chan distance, u graph.Node, v graph.Node, path Shortest, w float64)	{
	k := path.indexOf[v.ID()]
	j := path.indexOf[u.ID()]
	var changed bool
	joint := path.dist[j] + w
	if joint < path.dist[k] {
		changed = true
		fmt.Println(joint)
	}else{
		changed = false
	}
	chnl<-distance{to_idx: k, dist_new: joint, from_idx: j, changed: changed}
}
func BellmanFord(u graph.Node, g graph.Graph) (path Shortest) {
	if !g.Has(u) {
		return Shortest{from: u}
	}
	var weight Weighting
	if wg, ok := g.(graph.Weighter); ok {
		weight = wg.Weight
	} else {
		weight = UniformCost(g)
	}

	nodes := g.Nodes()

	path = newShortestFrom(u, nodes)
	path.dist[path.indexOf[u.ID()]] = 0

	chnl := make(chan distance)
	// TODO(kortschak): Consider adding further optimisations
	// from http://arxiv.org/abs/1111.5414.
	for i := 1; i < len(nodes); i++ {
		changed := false
		for _, u := range nodes {
			for _, v := range g.From(u) {
				w, ok := weight(u, v)
				if !ok {
					panic("bellman-ford: unexpected invalid weight")
				}
				if w<0{
					panic("bellman-ford: negative weight")
				}
				go UpdateDist(chnl, u, v, path, w)
			}
			for _, v := range g.From(u) {
				dist, ok := <- chnl
				fmt.Println(dist)
				if !ok{
					panic("bellman-ford: bad channel read")
					fmt.Println(v)
				}
				changed = dist.changed
				if(changed){
					path.set(dist.to_idx, dist.dist_new, dist.from_idx)	
				}
			}
		}
		if !changed {
			break
		}
	}

//	for j, u := range nodes {
//		for _, v := range g.From(u) {
//			k := path.indexOf[v.ID()]
//			w, ok := weight(u, v)
//			if !ok {
//				panic("bellman-ford: unexpected invalid weight")
//			}
///			if path.dist[j]+w < path.dist[k] {
	//			return path
	//		}
	//	}
//	}
	close(chnl)
	fmt.Println("BELLMAN FORD: ",path.dist)
	return path
}

// Apply the delta-stepping algorihtm to Graph and return
// a shortest path tree.
//
// Note that this uses Shortest to make it easier for you,
// but you can use another struct if that makes more sense
// for the concurrency model you chose.

func DeltaStep(s graph.Node, g graph.Graph) Shortest {
	delta := 3 //bin size
	var B []int[]graph.Node //sequence of buckets

	////////////////////
	for _, i := range B{
		S := {}
		for _, j := range B[i]{
			//req := getReqLight()
			S = append(S, B[i])
			for _, v in req{
				
			}
		}
	}
	//while B is not 0
		//S := {}
		//while B[i] is not {}
			//GET req req := {newdist+weight is in B[i] AND is light}
			//S := S append B[i]
			//for each v in req
				//relax v
			//for each v in req
				//colect to path
		//req := newdist+weight is in S AND is heavy
		//for each v in req
			//relax
	//i++ END while B is not 0

	//}
	return newShortestFrom(s, g.Nodes())
}
func relax(u graph.Node, v graph.Node, c int, path Shortest, chnl chan distance){
	from := path.indexOf[u.ID()]
	to := path.indexOf[v.ID()]

	if c < path.dist[path.indexOf(u)]{
		chnl<-distance{chnl<- distance{to_idx: to, dist_new: dist_updated, from_idx: from, changed: true}
		B[i].Push(v)//check this
	}else{
		chnl<-distance{chnl<- distance{to_idx: to, dist_new: dist_updated, from_idx: from, changed: false}
	}
}

// Runs dijkstra from gonum to make sure that the tests are correct.
func Dijkstra(s graph.Node, g graph.Graph) Shortest {
	return DijkstraFrom(s, g)
}
