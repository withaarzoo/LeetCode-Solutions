import (
	"sort"
)

type Point struct {
	pos int64
	x   int64
	y   int64
}

func maxDistance(side int, points [][]int, k int) int {
	s := int64(side)
	n := len(points)
	pts := make([]Point, n)

	for i, v := range points {
		x := int64(v[0])
		y := int64(v[1])

		var pos int64
		if x == 0 {
			pos = y
		} else if y == s {
			pos = s + x
		} else if x == s {
			pos = 3*s - y
		} else {
			pos = 4*s - x
		}
		pts[i] = Point{pos: pos, x: x, y: y}
	}

	sort.Slice(pts, func(i, j int) bool {
		return pts[i].pos < pts[j].pos
	})

	getOffset := func(x, y, d int64) int64 {
		if x == 0 {
			if d <= 2*s-y {
				return d
			}
			if d <= s+y {
				return 2*s + d - 2*y
			}
			return -1
		} else if y == s {
			if d <= 2*s-x {
				return d
			}
			if d <= s+x {
				return 2*s + d - 2*x
			}
			return -1
		} else if x == s {
			if d <= s+y {
				return d
			}
			if d <= 2*s-y {
				return d + 2*y
			}
			return -1
		} else {
			if d <= s+x {
				return d
			}
			if d <= 2*s-x {
				return d + 2*x
			}
			return -1
		}
	}

	lowerBound := func(arr []int64, l, r int, target int64) int {
		for l < r {
			m := (l + r) >> 1
			if arr[m] < target {
				l = m + 1
			} else {
				r = m
			}
		}
		return l
	}

	can := func(d int64) bool {
		pos3 := make([]int64, 3*n)
		for i := 0; i < n; i++ {
			pos3[i] = pts[i].pos
			pos3[i+n] = pts[i].pos + 4*s
			pos3[i+2*n] = pts[i].pos + 8*s
		}

		nxt := make([]int, 2*n)
		for i := range nxt {
			nxt[i] = -1
		}

		for i := 0; i < 2*n; i++ {
			p := pts[i%n]
			off := getOffset(p.x, p.y, d)
			if off < 0 {
				continue
			}

			target := pos3[i] + off
			hi := i + n
			if hi > 3*n {
				hi = 3*n
			}
			j := lowerBound(pos3, i+1, hi, target)
			if j < hi {
				nxt[i] = j
			}
		}

		for start := 0; start < n; start++ {
			cur := start
			cnt := 1

			for cnt < k {
				cur = nxt[cur]
				if cur == -1 || cur >= start+n {
					break
				}
				cnt++
			}

			if cnt >= k {
				a := pts[start]
				b := pts[cur%n]
				dist := a.x - b.x
				if dist < 0 {
					dist = -dist
				}
				dy := a.y - b.y
				if dy < 0 {
					dy = -dy
				}
				if dist+dy >= d {
					return true
				}
			}
		}

		return false
	}

	lo, hi := int64(0), 2*s
	for lo < hi {
		mid := (lo + hi + 1) >> 1
		if can(mid) {
			lo = mid
		} else {
			hi = mid - 1
		}
	}

	return int(lo)
}