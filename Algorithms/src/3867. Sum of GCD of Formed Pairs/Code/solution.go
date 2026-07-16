import "sort"

// Euclidean algorithm for gcd
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gcdSum(nums []int) int64 {

	n := len(nums)

	// Store prefix gcd values
	prefixGcd := make([]int, n)

	// Running prefix maximum
	prefixMax := 0

	// Build prefixGcd
	for i, x := range nums {
		if x > prefixMax {
			prefixMax = x
		}
		prefixGcd[i] = gcd(x, prefixMax)
	}

	// Sort the array
	sort.Ints(prefixGcd)

	var ans int64 = 0

	// Pair smallest with largest
	left := 0
	right := n - 1

	for left < right {
		ans += int64(gcd(prefixGcd[left], prefixGcd[right]))
		left++
		right--
	}

	return ans
}