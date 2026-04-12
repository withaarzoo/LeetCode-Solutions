func minimumDistance(word string) int {
    // Memoization map
    memo := make(map[[3]int]int)

    // Calculate Manhattan distance between two letters
    var getDist func(int, int) int
    getDist = func(a, b int) int {
        // 26 means finger is not placed yet
        if a == 26 || b == 26 {
            return 0
        }

        row1, col1 := a/6, a%6
        row2, col2 := b/6, b%6

        if row1 < row2 {
            row1, row2 = row2, row1
        }

        if col1 < col2 {
            col1, col2 = col2, col1
        }

        return (row1 - row2) + (col1 - col2)
    }

    var solve func(int, int, int) int
    solve = func(idx, f1, f2 int) int {
        // If all characters are typed
        if idx == len(word) {
            return 0
        }

        key := [3]int{idx, f1, f2}

        // Return memoized result
        if val, exists := memo[key]; exists {
            return val
        }

        cur := int(word[idx] - 'A')

        // Option 1: Use finger 1
        useFinger1 := getDist(f1, cur) + solve(idx+1, cur, f2)

        // Option 2: Use finger 2
        useFinger2 := getDist(f2, cur) + solve(idx+1, f1, cur)

        ans := useFinger1
        if useFinger2 < ans {
            ans = useFinger2
        }

        memo[key] = ans
        return ans
    }

    // Both fingers initially not placed
    return solve(0, 26, 26)
}