func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	n := len(source)

	parent := make([]int, n)
	rank := make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(a, b int) {
		pa := find(a)
		pb := find(b)

		if pa == pb {
			return
		}

		if rank[pa] < rank[pb] {
			parent[pa] = pb
		} else if rank[pb] < rank[pa] {
			parent[pb] = pa
		} else {
			parent[pb] = pa
			rank[pa]++
		}
	}

	// Build connected components
	for _, swap := range allowedSwaps {
		union(swap[0], swap[1])
	}

	// Group indices by root
	groups := make(map[int][]int)

	for i := 0; i < n; i++ {
		root := find(i)
		groups[root] = append(groups[root], i)
	}

	answer := 0

	// Process each component
	for _, indices := range groups {
		freq := make(map[int]int)

		// Count source values
		for _, idx := range indices {
			freq[source[idx]]++
		}

		// Match target values
		for _, idx := range indices {
			if freq[target[idx]] > 0 {
				freq[target[idx]]--
			} else {
				answer++
			}
		}
	}

	return answer
}