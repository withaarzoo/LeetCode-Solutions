class Solution {
    public int maximumLength(int[] nums) {
        // Store frequency of every number
        HashMap<Long, Integer> freq = new HashMap<>();

        for (int x : nums) {
            freq.put((long) x, freq.getOrDefault((long) x, 0) + 1);
        }

        int ans = 1;

        // Handle value 1 separately
        if (freq.containsKey(1L)) {
            int cnt = freq.get(1L);

            // Only odd count of ones is valid
            ans = Math.max(ans, (cnt % 2 == 1) ? cnt : cnt - 1);
        }

        // Try every distinct starting value
        for (long start : freq.keySet()) {
            if (start == 1L)
                continue;

            long cur = start;
            int len = 0;

            while (freq.containsKey(cur)) {
                // Use two copies if possible
                if (freq.get(cur) >= 2) {
                    len += 2;

                    // Move to the squared value
                    cur = cur * cur;
                } else {
                    // Single copy becomes the center
                    len++;
                    break;
                }
            }

            // No center found
            if ((len & 1) == 0)
                len--;

            ans = Math.max(ans, len);
        }

        return ans;
    }
}