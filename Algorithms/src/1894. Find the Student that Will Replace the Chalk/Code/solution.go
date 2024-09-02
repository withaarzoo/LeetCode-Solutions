func chalkReplacer(chalk []int, k int) int {
    // Step 1: Calculate the total amount of chalk needed for one complete round.
    // Initialize a variable to store the sum of chalk required by all students in one round.
    totalChalk := 0

    // Iterate through each student's chalk usage and sum them up.
    for _, c := range chalk {
        totalChalk += c
    }

    // Step 2: Reduce k by the total chalk used in one full round.
    // This is done using the modulo operation, which gives us the remainder after 
    // k is divided by totalChalk. Essentially, it reduces k to a smaller value that 
    // represents the remaining chalk after multiple full rounds.
    k %= totalChalk

    // Step 3: Iterate through the list of students to find out who will run out of chalk.
    // We check each student's chalk usage against the remaining chalk k.
    for i, c := range chalk {
        // If the remaining chalk (k) is less than the chalk required by the current student,
        // this student will be the one who needs to replace the chalk.
        if k < c {
            return i // Return the index of the student who will replace the chalk.
        }
        // Subtract the current student's chalk usage from k to check for the next student.
        k -= c
    }

    // Step 4: Safety return.
    // This line should never be reached due to the logic of the problem, 
    // but is included as a safeguard.
    return -1
}
