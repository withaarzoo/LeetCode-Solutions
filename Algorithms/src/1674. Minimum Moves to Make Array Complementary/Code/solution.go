func minMoves(nums []int, limit int) int {
    
    n := len(nums)

    // Difference array
    diff := make([]int, 2*limit+2)

    // Process every pair
    for i := 0; i < n/2; i++ {

        a := nums[i]
        b := nums[n-1-i]

        // Keep a <= b
        if a > b {
            a, b = b, a
        }

        // Range where 1 move is enough
        diff[a+1] -= 1
        diff[b+limit+1] += 1

        // Exact sum where 0 moves are needed
        diff[a+b] -= 1
        diff[a+b+1] += 1
    }

    pairs := n / 2

    // Initially assume 2 moves for every pair
    current := pairs * 2

    answer := current

    // Prefix sum traversal
    for sum := 2; sum <= 2*limit; sum++ {

        current += diff[sum]

        if current < answer {
            answer = current
        }
    }

    return answer
}