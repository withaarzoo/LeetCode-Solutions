func minMoves(classroom []string, energy int) int {
	m := len(classroom)
	n := len(classroom[0])

	// id[r][c] stores the bit assigned to a litter cell.
	id := make([][]int, m)
	for r := 0; r < m; r++ {
		id[r] = make([]int, n)

		// -1 means this cell does not contain litter.
		for c := 0; c < n; c++ {
			id[r][c] = -1
		}
	}

	k := 0
	sr, sc := 0, 0

	// Find S and assign a unique bit to every L cell.
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if classroom[r][c] == 'S' {
				sr, sc = r, c
			} else if classroom[r][c] == 'L' {
				id[r][c] = k
				k++
			}
		}
	}

	// If there is no litter, zero moves are required.
	if k == 0 {
		return 0
	}

	// This mask has every litter bit turned on.
	totalMask := (1 << k) - 1

	// best[r][c][mask] stores the maximum energy seen
	// for this position and collected-litter mask.
	best := make([][][]int, m)

	for r := 0; r < m; r++ {
		best[r] = make([][]int, n)

		for c := 0; c < n; c++ {
			best[r][c] = make([]int, 1<<k)

			// -1 means this position and mask has not been reached yet.
			for mask := 0; mask < (1 << k); mask++ {
				best[r][c][mask] = -1
			}
		}
	}

	// A state stores position, collected mask, remaining energy, and moves.
	type State struct {
		r, c   int
		mask   int
		energy int
		moves  int
	}

	// A slice with a head pointer works as an efficient BFS queue.
	queue := make([]State, 0)
	head := 0

	// Start from S with full energy and no collected litter.
	best[sr][sc][0] = energy
	queue = append(queue, State{sr, sc, 0, energy, 0})

	// Four possible movement directions.
	dr := [4]int{-1, 1, 0, 0}
	dc := [4]int{0, 0, -1, 1}

	for head < len(queue) {
		// Get the next state in BFS order.
		cur := queue[head]
		head++

		// Try moving in all four directions.
		for d := 0; d < 4; d++ {
			nr := cur.r + dr[d]
			nc := cur.c + dc[d]

			// Ignore positions outside the classroom.
			if nr < 0 || nr >= m || nc < 0 || nc >= n {
				continue
			}

			// Obstacles cannot be entered.
			if classroom[nr][nc] == 'X' {
				continue
			}

			// Every move consumes one unit of energy.
			ne := cur.energy - 1

			// The student cannot make a move without energy.
			if ne < 0 {
				continue
			}

			nmask := cur.mask

			// R restores the energy to its maximum capacity.
			if classroom[nr][nc] == 'R' {
				ne = energy
			}

			// Mark the litter as collected using its assigned bit.
			if classroom[nr][nc] == 'L' {
				nmask |= 1 << id[nr][nc]
			}

			// All litter has been collected, so return the shortest distance.
			if nmask == totalMask {
				return cur.moves + 1
			}

			// A state with less or equal energy is dominated by an existing one.
			if ne <= best[nr][nc][nmask] {
				continue
			}

			// Save the strongest energy value seen for this state.
			best[nr][nc][nmask] = ne

			// Add the improved state to the BFS queue.
			queue = append(queue, State{
				nr, nc, nmask, ne, cur.moves + 1,
			})
		}
	}

	// No valid path can collect every litter cell.
	return -1
}