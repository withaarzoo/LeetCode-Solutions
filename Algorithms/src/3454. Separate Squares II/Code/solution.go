func separateSquares(squares [][]int) float64 {
	type Event struct {
		y, x1, x2 float64
		t         int
	}

	var xs []float64
	var events []Event

	for _, s := range squares {
		x := float64(s[0])
		y := float64(s[1])
		l := float64(s[2])
		xs = append(xs, x, x+l)
		events = append(events,
			Event{y, x, x + l, 1},
			Event{y + l, x, x + l, -1},
		)
	}

	sort.Float64s(xs)
	xs = unique(xs)
	sort.Slice(events, func(i, j int) bool {
		return events[i].y < events[j].y
	})

	n := len(xs)
	cnt := make([]int, 4*n)
	seg := make([]float64, 4*n)

	var update func(int, int, int, int, int, int)
	update = func(node, l, r, ql, qr, val int) {
		if qr <= l || r <= ql {
			return
		}
		if ql <= l && r <= qr {
			cnt[node] += val
		} else {
			m := (l + r) >> 1
			update(node*2, l, m, ql, qr, val)
			update(node*2+1, m, r, ql, qr, val)
		}

		if cnt[node] > 0 {
			seg[node] = xs[r] - xs[l]
		} else if r-l == 1 {
			seg[node] = 0
		} else {
			seg[node] = seg[node*2] + seg[node*2+1]
		}
	}

	idx := map[float64]int{}
	for i, v := range xs {
		idx[v] = i
	}

	total := 0.0
	prevY := events[0].y
	strips := [][]float64{}

	for _, e := range events {
		if e.y > prevY {
			w := seg[1]
			h := e.y - prevY
			total += w * h
			strips = append(strips, []float64{prevY, h, w})
			prevY = e.y
		}
		update(1, 0, n-1, idx[e.x1], idx[e.x2], e.t)
	}

	half := total / 2
	acc := 0.0

	for _, s := range strips {
		if acc+s[1]*s[2] >= half {
			return s[0] + (half-acc)/s[2]
		}
		acc += s[1] * s[2]
	}
	return 0
}

func unique(arr []float64) []float64 {
	res := []float64{}
	for i, v := range arr {
		if i == 0 || v != arr[i-1] {
			res = append(res, v)
		}
	}
	return res
}
