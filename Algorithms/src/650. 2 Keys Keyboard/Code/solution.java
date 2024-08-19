class Solution {
    public int minSteps(int n) {
        // Initialize the number of operations required to 0
        int operations = 0;

        // Start checking for divisors from 2 up to n
        for (int i = 2; i <= n; i++) {
            // While 'n' is divisible by 'i', keep dividing and count the operations
            while (n % i == 0) {
                // Add 'i' to the operations count as each division represents 'i' copy-paste
                // operations
                operations += i;
                // Divide 'n' by 'i' to reduce 'n' for further factorization
                n /= i;
            }
        }

        // Return the total number of operations required to achieve 'n' 'A's
        return operations;
    }
}
