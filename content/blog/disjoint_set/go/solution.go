package disjoint_set

type DisjointSet struct {
	// Представление графа в виде слайса
	nodes []int
}

func NewDisjointSet(size int) DisjointSet {
	// Создаем слайс заданного размера по количеству элементов
	nodes := make([]int, size)

	for i := 0; i < size; i++ {
		// При инициализации считаем что каждый элемент обособлен и не соединен с другими.
		// В таком случае каждый элемент является рутовым сам для себя
		nodes[i] = i
	}

	return DisjointSet{
		nodes: nodes,
	}
}

func (uf *DisjointSet) find(node int) int {
	return uf.nodes[node]
}

func (uf *DisjointSet) connected(x int, y int) bool {
	return uf.find(x) == uf.find(y)
}

func (uf *DisjointSet) union(x int, y int) {
	rootX := uf.find(x)
	rootY := uf.find(y)

	// Производим объединение только если два элемента находятся не в одном множестве и не имеют одинаковый рутовый узел.
	if rootY != rootX {
		for i := range uf.nodes {
			// Обновляем значение рутового элемента на новое у всех элементов второго множества
			if uf.nodes[i] == rootY {
				uf.nodes[i] = rootX
			}
		}
	}
}
