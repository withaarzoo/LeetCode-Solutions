func twoEditWords(queries []string, dictionary []string) []string {
    result := []string{}

    // Check every query word
    for _, query := range queries {

        // Compare with every dictionary word
        for _, word := range dictionary {
            diff := 0

            // Count character differences
            for i := 0; i < len(query); i++ {
                if query[i] != word[i] {
                    diff++
                }

                // Stop early if more than 2 edits are needed
                if diff > 2 {
                    break
                }
            }

            // If query can match within 2 edits
            if diff <= 2 {
                result = append(result, query)
                break
            }
        }
    }

    return result
}