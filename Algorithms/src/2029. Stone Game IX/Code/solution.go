func stoneGameIX(stones []int) bool {
	// cnt[r] stores how many stones have remainder r modulo 3.
	cnt := [3]int{}

	// Count the stones in each remainder group.
	for _, stone := range stones {
		// Only the remainder matters for the game.
		cnt[stone%3]++
	}

	// With an even number of remainder-0 stones,
	// Alice needs both a remainder-1 and a remainder-2 stone.
	if cnt[0]%2 == 0 {
		return cnt[1] > 0 && cnt[2] > 0
	}

	// Calculate the difference between the two useful groups.
	diff := cnt[1] - cnt[2]

	// Make the difference positive.
	if diff < 0 {
		diff = -diff
	}

	// With an odd number of zero-remainder stones,
	// a difference greater than 2 gives Alice a winning strategy.
	return diff > 2
}