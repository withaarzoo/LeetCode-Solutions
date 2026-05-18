func minJumps(arr []int) int {

	n := len(arr)

	// No jump needed
	if n == 1 {
		return 0
	}

	// Store all indices for every value
	mp := make(map[int][]int)

	for i, val := range arr {
		mp[val] = append(mp[val], i)
	}

	// BFS queue
	queue := []int{0}

	// Visited array
	visited := make([]bool, n)

	visited[0] = true

	steps := 0

	for len(queue) > 0 {

		size := len(queue)

		// Process one BFS level
		for i := 0; i < size; i++ {

			idx := queue[0]
			queue = queue[1:]

			// Last index reached
			if idx == n-1 {
				return steps
			}

			// Move left
			if idx-1 >= 0 && !visited[idx-1] {
				visited[idx-1] = true
				queue = append(queue, idx-1)
			}

			// Move right
			if idx+1 < n && !visited[idx+1] {
				visited[idx+1] = true
				queue = append(queue, idx+1)
			}

			// Move to same-value indices
			for _, nextIdx := range mp[arr[idx]] {

				if !visited[nextIdx] {
					visited[nextIdx] = true
					queue = append(queue, nextIdx)
				}
			}

			// Clear processed group
			mp[arr[idx]] = []int{}
		}

		// One BFS level completed
		steps++
	}

	return -1
}