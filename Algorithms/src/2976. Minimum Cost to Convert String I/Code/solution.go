func minimumCost(source string, target string,
	original []byte, changed []byte, cost []int) int64 {

	const INF int64 = 1e18
	dist := make([][]int64, 26)
	for i := 0; i < 26; i++ {
		dist[i] = make([]int64, 26)
		for j := 0; j < 26; j++ {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = INF
			}
		}
	}

	for i := 0; i < len(original); i++ {
		u := original[i] - 'a'
		v := changed[i] - 'a'
		if int64(cost[i]) < dist[u][v] {
			dist[u][v] = int64(cost[i])
		}
	}

	// Floyd-Warshall
	for k := 0; k < 26; k++ {
		for i := 0; i < 26; i++ {
			for j := 0; j < 26; j++ {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	var ans int64 = 0
	for i := 0; i < len(source); i++ {
		s := source[i] - 'a'
		t := target[i] - 'a'
		if dist[s][t] == INF {
			return -1
		}
		ans += dist[s][t]
	}

	return ans
}
