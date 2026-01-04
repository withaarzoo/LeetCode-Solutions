func sumFourDivisors(nums []int) int {
    totalSum := 0

    for _, num := range nums {
        cnt := 0
        sum := 0

        for d := 1; d*d <= num; d++ {
            if num%d == 0 {
                other := num / d

                cnt++
                sum += d

                if other != d {
                    cnt++
                    sum += other
                }

                if cnt > 4 {
                    break
                }
            }
        }

        if cnt == 4 {
            totalSum += sum
        }
    }

    return totalSum
}
