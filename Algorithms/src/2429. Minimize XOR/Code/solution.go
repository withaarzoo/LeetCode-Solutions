func minimizeXor(num1 int, num2 int) int {
    count2 := bits.OnesCount(uint(num2)) // Number of 1s in num2
    count1 := bits.OnesCount(uint(num1)) // Number of 1s in num1

    if count1 == count2 {
        return num1
    }

    result := num1
    if count1 > count2 {
        for i := 0; i < 32 && count1 > count2; i++ {
            if result&(1<<i) != 0 {
                result &= ^(1 << i) // Clear bit
                count1--
            }
        }
    } else {
        for i := 0; i < 32 && count1 < count2; i++ {
            if result&(1<<i) == 0 {
                result |= (1 << i) // Set bit
                count1++
            }
        }
    }
    return result
}
