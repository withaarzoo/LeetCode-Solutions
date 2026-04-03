func maxWalls(robots []int, distance []int, walls []int) int {
	n := len(robots)

	type Robot struct {
		pos  int
		dist int
	}

	arr := make([]Robot, 0, n+1)
	for i := 0; i < n; i++ {
		arr = append(arr, Robot{robots[i], distance[i]})
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].pos < arr[j].pos
	})

	sort.Ints(walls)

	// Dummy robot
	arr = append(arr, Robot{int(1e9), 0})

	lowerBound := func(target int) int {
		left, right := 0, len(walls)

		for left < right {
			mid := left + (right-left)/2

			if walls[mid] < target {
				left = mid + 1
			} else {
				right = mid
			}
		}

		return left
	}

	upperBound := func(target int) int {
		left, right := 0, len(walls)

		for left < right {
			mid := left + (right-left)/2

			if walls[mid] <= target {
				left = mid + 1
			} else {
				right = mid
			}
		}

		return left
	}

	countWalls := func(leftRange int, rightRange int) int {
		if leftRange > rightRange {
			return 0
		}

		return upperBound(rightRange) - lowerBound(leftRange)
	}

	dp := make([][2]int, n)

	dp[0][0] = countWalls(arr[0].pos-arr[0].dist, arr[0].pos)

	firstRightEnd := arr[0].pos + arr[0].dist
	if n > 1 {
		firstRightEnd = min(firstRightEnd, arr[1].pos-1)
	}

	dp[0][1] = countWalls(arr[0].pos, firstRightEnd)

	for i := 1; i < n; i++ {
		pos := arr[i].pos
		dist := arr[i].dist

		// Shoot right
		rightEnd := min(pos+dist, arr[i+1].pos-1)
		rightWalls := countWalls(pos, rightEnd)

		dp[i][1] = max(dp[i-1][0], dp[i-1][1]) + rightWalls

		// Shoot left
		leftStart := max(pos-dist, arr[i-1].pos+1)
		leftWalls := countWalls(leftStart, pos)

		dp[i][0] = dp[i-1][0] + leftWalls

		prevRightEnd := min(arr[i-1].pos+arr[i-1].dist, pos-1)

		overlapStart := leftStart
		overlapEnd := min(prevRightEnd, pos-1)

		overlapWalls := countWalls(overlapStart, overlapEnd)

		dp[i][0] = max(dp[i][0], dp[i-1][1]+leftWalls-overlapWalls)
	}

	return max(dp[n-1][0], dp[n-1][1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}