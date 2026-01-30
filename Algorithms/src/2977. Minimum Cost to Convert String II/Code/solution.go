func minimumCost(source string, target string,
    original []string, changed []string, cost []int) int64 {

    const INF int64 = 1<<62
    id := map[string]int{}
    lens := map[int]bool{}
    sz := 0

    dist := make([][]int64, 201)
    for i := range dist {
        dist[i] = make([]int64, 201)
        for j := range dist[i] {
            dist[i][j] = INF
        }
    }

    for i := 0; i < len(original); i++ {
        if _, ok := id[original[i]]; !ok {
            id[original[i]] = sz
            lens[len(original[i])] = true
            sz++
        }
        if _, ok := id[changed[i]]; !ok {
            id[changed[i]] = sz
            sz++
        }
        u := id[original[i]]
        v := id[changed[i]]
        if int64(cost[i]) < dist[u][v] {
            dist[u][v] = int64(cost[i])
        }
    }

    for i := 0; i < sz; i++ {
        dist[i][i] = 0
    }

    for k := 0; k < sz; k++ {
        for i := 0; i < sz; i++ {
            if dist[i][k] < INF {
                for j := 0; j < sz; j++ {
                    if dist[k][j] < INF {
                        if dist[i][k]+dist[k][j] < dist[i][j] {
                            dist[i][j] = dist[i][k] + dist[k][j]
                        }
                    }
                }
            }
        }
    }

    n := len(source)
    dp := make([]int64, n+1)
    for i := range dp {
        dp[i] = INF
    }
    dp[0] = 0

    for i := 0; i < n; i++ {
        if dp[i] == INF {
            continue
        }
        if source[i] == target[i] {
            if dp[i] < dp[i+1] {
                dp[i+1] = dp[i]
            }
        }
        for L := range lens {
            if i+L > n {
                continue
            }
            s := source[i : i+L]
            t := target[i : i+L]
            if x, ok1 := id[s]; ok1 {
                if y, ok2 := id[t]; ok2 {
                    if dist[x][y] < INF {
                        if dp[i]+dist[x][y] < dp[i+L] {
                            dp[i+L] = dp[i] + dist[x][y]
                        }
                    }
                }
            }
        }
    }

    if dp[n] == INF {
        return -1
    }
    return dp[n]
}
