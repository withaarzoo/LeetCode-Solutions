func primeSubOperation(nums []int) bool {
	primes := generatePrimes(1000)
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			continue
		}
		
		target := nums[i] - nums[i+1]
		for _, p := range primes {
			if p > target {
				nums[i] -= p
				break
			}
		}
		if nums[i] <= 0 || nums[i] >= nums[i+1] {
			return false
		}
	}
	return true
}

func generatePrimes(maxNum int) []int {
	isPrime := make([]bool, maxNum+1)
	for i := 2; i <= maxNum; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= maxNum; i++ {
		if isPrime[i] {
			for j := i * i; j <= maxNum; j += i {
				isPrime[j] = false
			}
		}
	}
	var primes []int
	for i := 2; i <= maxNum; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	return primes
}