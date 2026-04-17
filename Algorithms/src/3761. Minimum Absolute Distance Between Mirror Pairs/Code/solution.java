class Solution {
    // Function to reverse digits of a number
    private int reverseNum(int x) {
        int rev = 0;

        while (x > 0) {
            rev = rev * 10 + (x % 10);
            x /= 10;
        }

        return rev;
    }

    public int minMirrorPairDistance(int[] nums) {
        HashMap<Integer, Integer> lastIndex = new HashMap<>();
        int ans = Integer.MAX_VALUE;

        for (int i = 0; i < nums.length; i++) {
            // If current number already exists in map,
            // then we found a mirror pair
            if (lastIndex.containsKey(nums[i])) {
                ans = Math.min(ans, i - lastIndex.get(nums[i]));
            }

            // Store reverse(nums[i]) with current index
            int rev = reverseNum(nums[i]);
            lastIndex.put(rev, i);
        }

        return ans == Integer.MAX_VALUE ? -1 : ans;
    }
}