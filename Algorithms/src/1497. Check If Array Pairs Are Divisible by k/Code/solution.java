class Solution {
    public boolean canArrange(int[] arr, int k) {
        // Frequency array to store the count of remainders
        int[] remainderFreq = new int[k];

        // Step 1: Calculate the remainder for each element and store the frequency
        for (int num : arr) {
            int remainder = (num % k + k) % k; // Ensure non-negative remainder
            remainderFreq[remainder]++;
        }

        // Step 2: Check if the pairing condition holds
        for (int i = 0; i <= k / 2; i++) {
            if (i == 0) {
                // Elements with remainder 0 must pair among themselves
                if (remainderFreq[i] % 2 != 0)
                    return false;
            } else {
                // Remainder i must pair with remainder k-i
                if (remainderFreq[i] != remainderFreq[k - i])
                    return false;
            }
        }

        return true;
    }
}