func largestAltitude(gain []int) int {
    // Current altitude starts at 0
    currentAltitude := 0

    // Highest altitude seen so far
    maxAltitude := 0

    // Process every gain value
    for _, change := range gain {
        // Apply altitude change
        currentAltitude += change

        // Update highest altitude if current altitude is greater
        if currentAltitude > maxAltitude {
            maxAltitude = currentAltitude
        }
    }

    // Return the highest altitude reached
    return maxAltitude
}