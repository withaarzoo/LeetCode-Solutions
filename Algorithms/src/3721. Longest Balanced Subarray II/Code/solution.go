// Go (clear and efficient)
package main
import (
    "fmt"
)

type SegTree struct {
    n int
    mn, mx, lazy []int
}

func NewSegTree(n int) *SegTree {
    return &SegTree{
        n: n,
        mn: make([]int, 4*n),
        mx: make([]int, 4*n),
        lazy: make([]int, 4*n),
    }
}
func (st *SegTree) apply(idx, v int) {
    st.mn[idx] += v
    st.mx[idx] += v
    st.lazy[idx] += v
}
func (st *SegTree) push(idx int) {
    z := st.lazy[idx]
    if z != 0 {
        st.apply(idx<<1, z)
        st.apply(idx<<1|1, z)
        st.lazy[idx] = 0
    }
}
func (st *SegTree) pull(idx int) {
    if st.mn[idx<<1] < st.mn[idx<<1|1] {
        st.mn[idx] = st.mn[idx<<1]
    } else {
        st.mn[idx] = st.mn[idx<<1|1]
    }
    if st.mx[idx<<1] > st.mx[idx<<1|1] {
        st.mx[idx] = st.mx[idx<<1]
    } else {
        st.mx[idx] = st.mx[idx<<1|1]
    }
}
func (st *SegTree) addRange(idx, l, r, ql, qr, val int) {
    if ql > qr { return }
    if ql <= l && r <= qr {
        st.apply(idx, val); return
    }
    st.push(idx)
    mid := (l + r) >> 1
    if ql <= mid { st.addRange(idx<<1, l, mid, ql, min(qr, mid), val) }
    if qr > mid  { st.addRange(idx<<1|1, mid+1, r, max(ql, mid+1), qr, val) }
    st.pull(idx)
}
func (st *SegTree) Add(l, r, v int) {
    if l > r { return }
    st.addRange(1, 0, st.n-1, l, r, v)
}
func (st *SegTree) findRightmostZero(idx, l, r, ql, qr int) int {
    if ql > qr || qr < l || ql > r { return -1 }
    if st.mn[idx] > 0 || st.mx[idx] < 0 { return -1 }
    if l == r {
        if st.mn[idx] == 0 { return l }
        return -1
    }
    st.push(idx)
    mid := (l + r) >> 1
    if qr > mid {
        res := st.findRightmostZero(idx<<1|1, mid+1, r, max(ql, mid+1), qr)
        if res != -1 { return res }
    }
    if ql <= mid {
        return st.findRightmostZero(idx<<1, l, mid, ql, min(qr, mid))
    }
    return -1
}
func (st *SegTree) FindRightmost(l, r int) int {
    if l > r { return -1 }
    return st.findRightmostZero(1, 0, st.n-1, l, r)
}
func min(a,b int) int { if a<b { return a }; return b }
func max(a,b int) int { if a>b { return a }; return b }

// Example wrapper function (LeetCode expects a function)
func longestBalanced(nums []int) int {
    n := len(nums)
    pos := map[int][]int{}
    for i, v := range nums {
        pos[v] = append(pos[v], i)
    }
    st := NewSegTree(n)
    for v, arr := range pos {
        sign := -1
        if v&1 == 1 { sign = 1 }
        st.Add(arr[0], n-1, sign)
    }
    ptr := map[int]int{}
    for k := range pos { ptr[k] = 0 }
    ans := 0
    for l := 0; l < n; l++ {
        r := st.FindRightmost(l, n-1)
        if r != -1 && r-l+1 > ans { ans = r-l+1 }
        x := nums[l]
        pIndex := ptr[x]; ptr[x] = pIndex+1
        arr := pos[x]
        nextPos := n
        if ptr[x] < len(arr) { nextPos = arr[ptr[x]] }
        sign := -1
        if x&1 == 1 { sign = 1 }
        L, R := l, nextPos-1
        if L <= R { st.Add(L, R, -sign) }
    }
    return ans
}

// main is just for quick local test (remove on leetcode)
func main() {
    fmt.Println(longestBalanced([]int{2,5,4,3})) // expect 4
}
