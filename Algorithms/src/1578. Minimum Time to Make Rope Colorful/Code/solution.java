class Solution {
    public int minCost(String colors, int[] neededTime) {
        long ans = 0L; // accumulate answer here (long for safety)
        long blockSum = 0L; // sum of times in current block
        int blockMax = 0; // maximum time in current block
        int n = colors.length();

        for (int i = 0; i < n; ++i) {
            if (i > 0 && colors.charAt(i) != colors.charAt(i - 1)) {
                ans += blockSum - blockMax;
                blockSum = 0;
                blockMax = 0;
            }
            blockSum += neededTime[i];
            blockMax = Math.max(blockMax, neededTime[i]);
        }
        ans += blockSum - blockMax;
        return (int) ans;
    }
}
