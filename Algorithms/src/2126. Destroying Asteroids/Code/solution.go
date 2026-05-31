func asteroidsDestroyed(mass int, asteroids []int) bool {

    // Sort asteroids from smallest to largest
    sort.Ints(asteroids)

    // Use int64 because mass can grow significantly
    currentMass := int64(mass)

    // Process each asteroid
    for _, asteroid := range asteroids {

        // Planet cannot destroy this asteroid
        if currentMass < int64(asteroid) {
            return false
        }

        // Gain asteroid mass
        currentMass += int64(asteroid)
    }

    // All asteroids destroyed
    return true
}