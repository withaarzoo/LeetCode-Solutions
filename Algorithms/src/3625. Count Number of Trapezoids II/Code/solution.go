package main

func countTrapezoids(points [][]int) int {
	n := len(points)
	const SHIFT = 3000

	encodePair := func(a, b int) int {
		return ((a + SHIFT) << 13) ^ (b + SHIFT)
	}

	gcd := func(a, b int) int {
		if a < 0 {
			a = -a
		}
		if b < 0 {
			b = -b
		}
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	// outer key -> inner map(lineId -> count)
	bySlope := make(map[int]map[int]int)
	byVector := make(map[int]map[int]int)

	addTo := func(mp map[int]map[int]int, key, lineId int) {
		if mp[key] == nil {
			mp[key] = make(map[int]int)
		}
		mp[key][lineId]++
	}

	for i := 0; i < n; i++ {
		x1, y1 := points[i][0], points[i][1]
		for j := i + 1; j < n; j++ {
			x2, y2 := points[j][0], points[j][1]
			dx := x2 - x1
			dy := y2 - y1

			if dx < 0 || (dx == 0 && dy < 0) {
				dx = -dx
				dy = -dy
			}

			g := gcd(dx, dy)
			ux := dx / g
			uy := dy / g

			lineId := ux*y1 - uy*x1

			slopeKey := encodePair(ux, uy)
			vectorKey := encodePair(dx, dy)

			addTo(bySlope, slopeKey, lineId)
			addTo(byVector, vectorKey, lineId)
		}
	}

	countPairs := func(mp map[int]map[int]int) int64 {
		var res int64
		for _, inner := range mp {
			var sum, sumSq int64
			for _, c := range inner {
				cc := int64(c)
				sum += cc
				sumSq += cc * cc
			}
			res += (sum*sum - sumSq) / 2
		}
		return res
	}

	withParallel := countPairs(bySlope)
	parallelogramTwice := countPairs(byVector)

	ans := withParallel - parallelogramTwice/2
	return int(ans)
}
