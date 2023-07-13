package algorithm

type Graph struct {
	Nodes 	 []string				`json:"nodes"`
	Edges    map[string][]string	`json:"edges"`
}

// Helper function to get the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TarjanSCC(g *Graph) [][]string {
	// Initialize variables
	index := 0
	stack := make([]string, 0)
	indices := make(map[string]int)
	lowLinks := make(map[string]int)
	onStack := make(map[string]bool)
	result := make([][]string, 0)

	// Recursive depth-first search function
	var strongConnect func(v string)

	strongConnect = func(v string) {
		// Set the depth index and low link value of the current vertex
		indices[v] = index
		lowLinks[v] = index
		index++
		stack = append(stack, v)
		onStack[v] = true

		// Consider all neighbors of the current vertex
		for _, w := range g.Edges[v] {
			if indices[w] == 0 {
				// Neighbor w has not yet been visited, so recursively visit it
				strongConnect(w)
				lowLinks[v] = min(lowLinks[v], lowLinks[w])
			} else if onStack[w] {
				// Neighbor w is on the stack and hence in the current SCC
				lowLinks[v] = min(lowLinks[v], indices[w])
			}
		}

		// If v is a root node, pop the stack and generate an SCC
		if lowLinks[v] == indices[v] {
			scc := make([]string, 0)
			w := ""
			for w != v {
				w = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				onStack[w] = false
				scc = append(scc, w)
			}
			result = append(result, scc)
		}
	}

	// Perform Tarjan's algorithm on each unvisited vertex
	for _, v := range g.Nodes {
		if indices[v] == 0 {
			strongConnect(v)
		}
	}

	return result
}
