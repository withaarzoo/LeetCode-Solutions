class Solution { 
    public long[] resultArray(int[] nums, int k) { 
        long[] result = new long[k];
        long[] curr = new long[k];
        for (int num : nums) {
            int m = num % k;
            long[] next = new long[k];
            next[m] = 1;
            for (int prev = 0; prev < k; prev++) {
                if (curr[prev] > 0) {
                    int nr = (int)((prev * 1L * m) % k);
                    next[nr] += curr[prev];
                }
            }
            for (int r = 0; r < k; r++) {
                result[r] += next[r];
            }
            curr = next;
        }
        return result;
    } 
}