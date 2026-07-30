func minimumPushes(word string) int {

    // Store the final answer
    pushes := 0

    // Traverse every character
    for i := 0; i < len(word); i++ {

        // Every block of 8 letters increases the cost by 1
        pushes += (i / 8) + 1
    }

    // Return the minimum number of pushes
    return pushes
}