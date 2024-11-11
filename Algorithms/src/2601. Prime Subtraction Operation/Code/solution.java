class Solution {
    private boolean isPrime(int number) {
        if (number < 2)
            return false;

        for (int i = 2; i <= Math.sqrt(number); i++) {
            if (number % i == 0) {
                return false;
            }
        }
        return true;
    }

    public boolean primeSubOperation(int[] nums) {
        for (int i = 0; i < nums.length; i++) {
            int bound = (i == 0) ? nums[0] : nums[i] - nums[i - 1];

            if (bound <= 0) {
                return false;
            }

            int largestPrime = 0;
            for (int j = bound - 1; j >= 2; j--) {
                if (isPrime(j)) {
                    largestPrime = j;
                    break;
                }
            }

            nums[i] -= largestPrime;
        }

        return true;
    }
}