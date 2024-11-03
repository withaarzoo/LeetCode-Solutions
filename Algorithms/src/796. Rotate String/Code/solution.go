func rotateString(s string, goal string) bool {
    // If lengths differ, they cannot be rotations
    if len(s) != len(goal) {
        return false
    }
    
    // Concatenate s with itself
    doubled := s + s
    
    // Check if goal is a substring of doubled
    return strings.Contains(doubled, goal)
}
