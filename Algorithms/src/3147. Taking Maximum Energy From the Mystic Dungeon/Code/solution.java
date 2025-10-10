class Solution {
    public int maximumEnergy(int[] energy, int k) {
        int n = energy.length;
        long ans = Long.MIN_VALUE;
        for (int r = 0; r < k; ++r) {
            long cur = 0;
            int last = r + ((n - 1 - r) / k) * k; // last index in this class
            for (int i = last; i >= r; i -= k) {
                cur += energy[i];     // suffix sum from i to end of class
                ans = Math.max(ans, cur);
            }
        }
        return (int) ans;
    }
}
