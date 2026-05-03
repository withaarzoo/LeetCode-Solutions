func rotateString(s string, goal string) bool {
    // If lengths are different, rotation is impossible
    if len(s) != len(goal) {
        return false
    }

    // Concatenate s with itself
    doubled := s + s

    // Check if goal exists in doubled string
    return strings.Contains(doubled, goal)
}