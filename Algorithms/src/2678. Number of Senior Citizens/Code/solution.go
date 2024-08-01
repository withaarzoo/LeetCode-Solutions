// countSeniors calculates the number of seniors (age > 60) in the provided details slice.
func countSeniors(details []string) int {
    // Initialize the count of seniors
    count := 0
    
    // Iterate over each detail string in the slice
    for _, detail := range details {
        // Extract the age substring from the string (positions 11 to 12, inclusive)
        ageStr := detail[11:13]
        
        // Convert the age substring to an integer
        age, _ := strconv.Atoi(ageStr)
        
        // If the converted age is greater than 60, increment the count
        if age > 60 {
            count++
        }
    }
    
    // Return the total count of seniors
    return count
}