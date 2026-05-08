func minJumps(nums []int) int {

	n := len(nums)

	// Already at destination
	if n == 1 {
		return 0
	}

	// Find maximum value
	mx := 0

	for _, x := range nums {
		if x > mx {
			mx = x
		}
	}

	// Smallest prime factor array
	spf := make([]int, mx+1)

	// Initialize SPF
	for i := 0; i <= mx; i++ {
		spf[i] = i
	}

	// Build sieve
	for i := 2; i*i <= mx; i++ {

		if spf[i] == i {

			for j := i * i; j <= mx; j += i {

				if spf[j] == j {
					spf[j] = i
				}
			}
		}
	}

	// Prime factor -> indices mapping
	mp := map[int][]int{}

	for i, val := range nums {

		x := val

		used := map[int]bool{}

		// Extract unique prime factors
		for x > 1 {

			p := spf[x]

			if !used[p] {

				mp[p] = append(mp[p], i)

				used[p] = true
			}

			x /= p
		}
	}

	// BFS queue
	q := []int{0}

	// Distance array
	dist := make([]int, n)

	for i := 0; i < n; i++ {
		dist[i] = -1
	}

	dist[0] = 0

	front := 0

	for front < len(q) {

		i := q[front]
		front++

		steps := dist[i]

		// Reached destination
		if i == n-1 {
			return steps
		}

		// Move left
		if i-1 >= 0 && dist[i-1] == -1 {

			dist[i-1] = steps + 1

			q = append(q, i-1)
		}

		// Move right
		if i+1 < n && dist[i+1] == -1 {

			dist[i+1] = steps + 1

			q = append(q, i+1)
		}

		val := nums[i]

		// Teleport allowed only if value is prime
		if val > 1 && spf[val] == val {

			for _, nxt := range mp[val] {

				if dist[nxt] == -1 {

					dist[nxt] = steps + 1

					q = append(q, nxt)
				}
			}

			// Clear after use
			mp[val] = []int{}
		}
	}

	return -1
}