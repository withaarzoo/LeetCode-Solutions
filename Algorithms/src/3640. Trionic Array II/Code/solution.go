func maxSumTrionic(nums []int) int64 {
    n := len(nums)
    left := make([]int64, n)
    right := make([]int64, n)

    for i := 0; i < n; i++ {
        left[i] = int64(nums[i])
        if i > 0 && nums[i-1] < nums[i] && left[i-1] > 0 {
            left[i] += left[i-1]
        }
    }

    for i := n-1; i >= 0; i-- {
        right[i] = int64(nums[i])
        if i+1 < n && nums[i] < nums[i+1] && right[i+1] > 0 {
            right[i] += right[i+1]
        }
    }

    type Block struct {
        l, r int
        sum int64
    }

    blocks := []Block{}
    l, s := 0, int64(nums[0])
    for i := 1; i < n; i++ {
        if nums[i-1] <= nums[i] {
            blocks = append(blocks, Block{l, i-1, s})
            l = i
            s = 0
        }
        s += int64(nums[i])
    }
    blocks = append(blocks, Block{l, n-1, s})

    ans := int64(-1e18)
    for _, b := range blocks {
        if b.l > 0 && b.r < n-1 &&
           nums[b.l-1] < nums[b.l] &&
           nums[b.r] < nums[b.r+1] &&
           b.l < b.r {
            val := left[b.l-1] + b.sum + right[b.r+1]
            if val > ans {
                ans = val
            }
        }
    }
    return ans
}
