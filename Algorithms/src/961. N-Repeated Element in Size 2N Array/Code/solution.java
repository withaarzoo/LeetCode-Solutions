class Solution {
    public int repeatedNTimes(int[] nums) {
        HashSet<Integer> seen = new HashSet<>();

        for (int x : nums) {
            // If number already exists, return it
            if (seen.contains(x)) {
                return x;
            }
            // Otherwise, store it
            seen.add(x);
        }
        return -1; // Guaranteed answer exists
    }
}
