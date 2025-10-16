class Solution {
    public int findSmallestInteger(int[] nums, int value) {
        int[] freq = new int[value];
        for (int a : nums) {
            int r = ((a % value) + value) % value; // handle negatives
            freq[r]++;
        }
        int x = 0;
        while (true) {
            int r = x % value;
            if (freq[r] == 0) return x;
            freq[r]--;
            x++;
        }
    }
}
