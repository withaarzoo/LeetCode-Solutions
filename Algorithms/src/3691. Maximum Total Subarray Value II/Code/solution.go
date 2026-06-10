func maxTotalValue(nums []int, k int) int64 {
	n := len(nums)

	// floor(log2(i))
	lg := make([]int, n+1)
	for i := 2; i <= n; i++ {
		lg[i] = lg[i/2] + 1
	}

	K := lg[n] + 1

	// Sparse table for maximums
	mx := make([][]int, K)

	// Sparse table for minimums
	mn := make([][]int, K)

	for i := 0; i < K; i++ {
		mx[i] = make([]int, n)
		mn[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		mx[0][i] = nums[i]
		mn[0][i] = nums[i]
	}

	// Build sparse tables
	for j := 1; j < K; j++ {
		for i := 0; i+(1<<j) <= n; i++ {
			a := mx[j-1][i]
			b := mx[j-1][i+(1<<(j-1))]
			if b > a {
				a = b
			}
			mx[j][i] = a

			c := mn[j-1][i]
			d := mn[j-1][i+(1<<(j-1))]
			if d < c {
				c = d
			}
			mn[j][i] = c
		}
	}

	// O(1) range value query
	getValue := func(l, r int) int64 {
		length := r - l + 1
		p := lg[length]

		a := mx[p][l]
		b := mx[p][r-(1<<p)+1]
		if b > a {
			a = b
		}

		c := mn[p][l]
		d := mn[p][r-(1<<p)+1]
		if d < c {
			c = d
		}

		return int64(a - c)
	}

	type Node struct {
		val int64
		l   int
		r   int
	}

	h := &MaxHeap{}
	heap.Init(h)

	for l := 0; l < n; l++ {
		heap.Push(h, Node{
			val: getValue(l, n-1),
			l:   l,
			r:   n - 1,
		})
	}

	var ans int64 = 0

	for k > 0 {
		cur := heap.Pop(h).(Node)

		ans += cur.val

		if cur.r > cur.l {
			heap.Push(h, Node{
				val: getValue(cur.l, cur.r-1),
				l:   cur.l,
				r:   cur.r - 1,
			})
		}

		k--
	}

	return ans
}

type MaxHeap []struct {
	val int64
	l   int
	r   int
}

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].val > h[j].val
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(struct {
		val int64
		l   int
		r   int
	}))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)

	item := old[n-1]
	*h = old[:n-1]

	return item
}