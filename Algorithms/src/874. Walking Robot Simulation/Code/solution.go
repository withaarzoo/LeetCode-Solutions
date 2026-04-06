func robotSim(commands []int, obstacles [][]int) int {
    // Store obstacles in a map for O(1) lookup
    obstacleSet := make(map[string]bool)

    for _, obs := range obstacles {
        key := fmt.Sprintf("%d,%d", obs[0], obs[1])
        obstacleSet[key] = true
    }

    // Directions: North, East, South, West
    dx := []int{0, 1, 0, -1}
    dy := []int{1, 0, -1, 0}

    dir := 0 // Start facing North
    x, y := 0, 0
    maxDistance := 0

    for _, command := range commands {
        // Turn right
        if command == -1 {
            dir = (dir + 1) % 4
        } else if command == -2 {
            // Turn left
            dir = (dir + 3) % 4
        } else {
            // Move forward step by step
            for step := 0; step < command; step++ {
                nextX := x + dx[dir]
                nextY := y + dy[dir]

                key := fmt.Sprintf("%d,%d", nextX, nextY)

                // Stop if obstacle exists
                if obstacleSet[key] {
                    break
                }

                x = nextX
                y = nextY

                distance := x*x + y*y
                if distance > maxDistance {
                    maxDistance = distance
                }
            }
        }
    }

    return maxDistance
}