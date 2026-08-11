func missingInteger(nums []int) int {
    // I start the sum with the first element because
    // nums[0] is always part of the sequential prefix.
    sum := nums[0]

    // I scan from the second element to find the end of the sequential prefix.
    for i := 1; i < len(nums); i++ {
        // The sequence continues only when the current value
        // is exactly one greater than the previous value.
        if nums[i] == nums[i-1]+1 {
            sum += nums[i]
        } else {
            // The sequence breaks here, so I stop the prefix scan.
            break
        }
    }

    // I use a map as a set to store every value from the array.
    seen := make(map[int]bool)

    // I mark every array value as present.
    for _, num := range nums {
        seen[num] = true
    }

    // I start checking candidates from the sequential prefix sum.
    answer := sum

    // If the candidate exists, I keep increasing it until
    // I find a number that is missing from the array.
    for seen[answer] {
        answer++
    }

    // The first missing candidate is the required answer.
    return answer
}