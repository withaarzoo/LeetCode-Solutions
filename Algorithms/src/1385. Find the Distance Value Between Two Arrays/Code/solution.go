func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
    sort.Ints(arr2)
    ans := 0
    for _, x := range arr1 {
        left, right := 0, len(arr2)
        for left < right {
            mid := (left + right) / 2
            if arr2[mid] < x {
                left = mid + 1
            } else {
                right = mid
            }
        }
        ok := true
        if left < len(arr2) && abs(arr2[left]-x) <= d {
            ok = false
        }
        if left > 0 && abs(arr2[left-1]-x) <= d {
            ok = false
        }
        if ok {
            ans++
        }
    }
    return ans
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}