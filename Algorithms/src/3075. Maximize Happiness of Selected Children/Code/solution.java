class Solution {
    public long maximumHappinessSum(int[] happiness, int k) {
        // Sort the array
        Arrays.sort(happiness);

        long ans = 0;
        int n = happiness.length;

        // Pick from the largest values
        for (int i = 0; i < k; i++) {
            long curr = happiness[n - 1 - i] - i;
            if (curr > 0) {
                ans += curr;
            }
        }

        return ans;
    }
}
