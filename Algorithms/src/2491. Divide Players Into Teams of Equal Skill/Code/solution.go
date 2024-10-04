import "sort"

func dividePlayers(skill []int) int64 {
    // Step 1: Sort the skill array
    sort.Ints(skill)
    
    totalSkill := skill[0] + skill[len(skill)-1] // Required sum for each pair
    var chemistrySum int64 = 0

    // Step 2: Pair players using two pointers
    for i := 0; i < len(skill)/2; i++ {
        // Check if the sum of current pair matches the required totalSkill
        if skill[i]+skill[len(skill)-i-1] != totalSkill {
            return -1 // Invalid configuration, return -1
        }
        // Calculate the chemistry (product of pair) and add it to the sum
        chemistrySum += int64(skill[i]) * int64(skill[len(skill)-i-1])
    }

    return chemistrySum // Return total chemistry
}