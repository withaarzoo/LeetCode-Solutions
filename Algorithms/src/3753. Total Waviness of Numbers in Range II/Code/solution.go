func totalWaviness(num1 int64, num2 int64) int64 {

	type Node struct {
		cnt int64
		wav int64
	}

	var solve func(int64) int64

	solve = func(n int64) int64 {
		if n < 0 {
			return 0
		}

		s := []byte(fmt.Sprintf("%d", n))

		type State struct {
			pos        int
			started    int
			last       int
			secondLast int
		}

		memo := map[State]Node{}

		var dfs func(int, int, int, int, bool) Node

		dfs = func(pos, started, last, secondLast int, tight bool) Node {
			if pos == len(s) {
				return Node{1, 0}
			}

			st := State{pos, started, last, secondLast}

			if !tight {
				if val, ok := memo[st]; ok {
					return val
				}
			}

			limit := 9
			if tight {
				limit = int(s[pos] - '0')
			}

			res := Node{}

			for d := 0; d <= limit; d++ {
				ntight := tight && d == limit

				if started == 0 && d == 0 {
					nxt := dfs(pos+1, 0, 10, 10, ntight)

					res.cnt += nxt.cnt
					res.wav += nxt.wav
				} else {
					var add int64 = 0

					if started == 1 && secondLast != 10 {
						if (last > secondLast && last > d) ||
							(last < secondLast && last < d) {
							add = 1
						}
					}

					nSecondLast := 10
					if started == 1 {
						nSecondLast = last
					}

					nxt := dfs(pos+1, 1, d, nSecondLast, ntight)

					res.cnt += nxt.cnt
					res.wav += nxt.wav + add*nxt.cnt
				}
			}

			if !tight {
				memo[st] = res
			}

			return res
		}

		return dfs(0, 0, 10, 10, true).wav
	}

	return solve(num2) - solve(num1-1)
}