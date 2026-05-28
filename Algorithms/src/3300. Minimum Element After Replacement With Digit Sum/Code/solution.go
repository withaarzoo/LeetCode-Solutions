func minElement(nums []int) int {

    // Helper function to calculate digit sum
    digitSum := func(num int) int {
        sum := 0

        // Process every digit
        for num > 0 {
            sum += num % 10 // Add last digit
            num /= 10       // Remove last digit
        }

        return sum
    }

    ans := int(^uint(0) >> 1) // Maximum int value

    // Check digit sum of every number
    for _, num := range nums {
        current := digitSum(num)

        if current < ans {
            ans = current
        }
    }

    return ans
}