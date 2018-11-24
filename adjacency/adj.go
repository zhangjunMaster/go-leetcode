package main

//基于二维数组而建立的

type v string

type Adj struct {
	Nodes    []v
	Edges    [][]int32
	NumEdges int
	NumNodes int
}

type edge struct {
	v0     int32
	v1     int32
	weight int32
}

func NewAdj(nodes []v, edegs []edge) *Adj {
	Edges := make([][]int32, 0)
	adj := Adj{Nodes: nodes, NumEdges: len(nodes), Edges: Edges}
	for i, _ := range nodes {
		for j, _ := range nodes {
			if i == j {
				adj.Edges[i][j] = 0
			} else {
				adj.Edges[i][j] = INFINITY
			}
		}
	}

	for _, v := range edegs {
		adj.Edges[v.v0][v.v1] = v.weight
	}
	return &adj
}

/**
  //     A   B   C   D
  //A [  0   5   3   6]
  //B [  5   0   7   0]
  //C [  3   7   0   9]
  //D [  6   0   9   0]
*/
