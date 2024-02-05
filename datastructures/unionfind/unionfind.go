package unionfind

type UnionFind struct {
	parents []int
	rank    []int
}

func New(size int) *UnionFind {
	parents := make([]int, size)
	rank := make([]int, size)
	for i := 0; i < size; i++ {
		parents[i] = i
	}
	return &UnionFind{
		parents: parents,
		rank:    rank,
	}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parents[x] != x {
		return uf.Find(uf.parents[x])
	}
	return x
}

// Union defines union method for UnionFind data structure, attaching the smaller tree to the larger based on rank
func (uf *UnionFind) Union(x, y int) {
	rootX, rootY := uf.Find(x), uf.Find(y)

	if uf.rank[rootX] > uf.rank[rootY] {
		uf.parents[rootY] = rootX
	} else if uf.rank[rootY] > uf.rank[rootX] {
		uf.parents[rootX] = rootY
	} else {
		// equal ranks
		uf.parents[rootY] = rootX
		uf.rank[rootX]++
	}

}
