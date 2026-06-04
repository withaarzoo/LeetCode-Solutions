func totalWaviness(num1 int, num2 int) int {
    answer := 0

    // Check every number in the range
    for num := num1; num <= num2; num++ {
        s := strconv.Itoa(num)

        // Numbers with fewer than 3 digits have waviness 0
        if len(s) < 3 {
            continue
        }

        // Check every middle digit
        for i := 1; i < len(s)-1; i++ {
            // Peak condition
            if s[i] > s[i-1] && s[i] > s[i+1] {
                answer++
            } else if s[i] < s[i-1] && s[i] < s[i+1] {
                // Valley condition
                answer++
            }
        }
    }

    return answer
}