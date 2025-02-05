package number_of_provinces

type DisjointSet struct {
	nodes []int
}

func NewDisjointSet(size int) DisjointSet {
	nodes := make([]int, size)

	for i := 0; i < size; i++ {
		nodes[i] = i
	}

	return DisjointSet{
		nodes: nodes,
	}
}

func (uf *DisjointSet) find(node int) int {
	return uf.nodes[node]
}

func (uf *DisjointSet) union(x int, y int) {
	rootX := uf.find(x)
	rootY := uf.find(y)

	if rootY != rootX {
		for i := range uf.nodes {
			if uf.nodes[i] == rootY {
				uf.nodes[i] = rootX
			}
		}
	}
}

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	disjointSet := NewDisjointSet(n)

	numberOfProvinces := n

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 && disjointSet.find(i) != disjointSet.find(j) {
				numberOfProvinces -= 1
				disjointSet.union(i, j)
			}
		}
	}

	return numberOfProvinces
}
