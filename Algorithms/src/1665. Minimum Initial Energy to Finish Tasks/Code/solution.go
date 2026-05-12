func minimumEffort(tasks [][]int) int {

    // Sort by (minimum - actual) in descending order
    sort.Slice(tasks, func(i, j int) bool {
        return (tasks[i][1] - tasks[i][0]) > (tasks[j][1] - tasks[j][0])
    })

    answer := 0 // Minimum initial energy
    energy := 0 // Current energy

    // Process all tasks
    for _, task := range tasks {

        actual := task[0]
        minimum := task[1]

        // If energy is not enough,
        // increase it
        if energy < minimum {

            need := minimum - energy

            answer += need
            energy += need
        }

        // Complete the task
        energy -= actual
    }

    return answer
}