import "sort"

func lexicographicallySmallestArray(nums []int, limit int) []int {
	n := len(nums) // Store the size of the input array.

	// Each pair stores a value and its original index.
	type Pair struct {
		value int
		index int
	}

	// Create an array containing all values with their original positions.
	elements := make([]Pair, n)
	for i := 0; i < n; i++ {
		elements[i] = Pair{nums[i], i}
	}

	// Sort by value so connected values appear next to each other.
	sort.Slice(elements, func(i, j int) bool {
		return elements[i].value < elements[j].value
	})

	// Store the lexicographically smallest result.
	answer := make([]int, n)

	start := 0 // Marks the beginning of the current connected group.

	for start < n {
		end := start // Expand the current connected group.

		// Keep consecutive values together while their difference is
		// at most limit, meaning they are connected through swaps.
		for end+1 < n &&
			int64(elements[end+1].value)-int64(elements[end].value) <= int64(limit) {
			end++
		}

		// Collect all original indices belonging to this group.
		indices := make([]int, 0, end-start+1)
		for i := start; i <= end; i++ {
			indices = append(indices, elements[i].index)
		}

		// Sort positions so smaller values can be placed first.
		sort.Ints(indices)

		// Values in elements[start:end+1] are already sorted.
		for i := 0; i < len(indices); i++ {
			answer[indices[i]] = elements[start+i].value
		}

		// Continue with the next connected group.
		start = end + 1
	}

	return answer // Return the lexicographically smallest arrangement.
}