func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
    n := len(nums1)
    d := make([]int, n)
    var total int64
    mx := 0
    for i := 0; i < n; i++ {
        if nums1[i] > nums2[i] {
            d[i] = nums1[i] - nums2[i]
        } else {
            d[i] = nums2[i] - nums1[i]
        }
        total += int64(d[i])
        if d[i] > mx {
            mx = d[i]
        }
    }
    k := int64(k1) + int64(k2)
    if total <= k {
        return 0
    }
    left, right := 0, mx
    for left < right {
        mid := left + (right-left)/2
        var need int64
        for _, v := range d {
            if v > mid {
                need += int64(v - mid)
            }
        }
        if need <= k {
            right = mid
        } else {
            left = mid + 1
        }
    }
    for i := 0; i < n; i++ {
        if d[i] > left {
            k -= int64(d[i] - left)
        }
        if d[i] > left {
            d[i] = left
        }
    }
    for i := 0; i < n && k > 0; i++ {
        if d[i] == left {
            d[i]--
            k--
        }
    }
    var ans int64
    for _, v := range d {
        ans += int64(v) * int64(v)
    }
    return ans
}