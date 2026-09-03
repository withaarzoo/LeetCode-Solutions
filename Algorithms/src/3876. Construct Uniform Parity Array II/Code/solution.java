class Solution {
    public boolean uniformArray(int[] nums1) {
        // Use a large value as "not found yet" for both parities.
        int minOdd = Integer.MAX_VALUE;
        int minEven = Integer.MAX_VALUE;

        // Scan the array once to find the smallest odd and even values.
        for (int x : nums1) {
            if (x % 2 == 0) {
                // Keep the smallest even number because it is the hardest one to change.
                minEven = Math.min(minEven, x);
            } else {
                // Keep the smallest odd number because it is the best number to subtract.
                minOdd = Math.min(minOdd, x);
            }
        }

        // If there is no odd number, all elements are already even.
        if (minOdd == Integer.MAX_VALUE) {
            return true;
        }

        // Every even value must be greater than the smallest odd value.
        // Subtracting that odd value then makes every even value odd.
        return minOdd < minEven;
    }
}