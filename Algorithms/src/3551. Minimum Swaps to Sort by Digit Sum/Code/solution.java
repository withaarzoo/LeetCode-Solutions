class Solution {
    public int minSwaps(int[] nums) {
        int n = nums.length;
        Integer[] sorted = new Integer[n];
        for (int i = 0; i < n; i++) sorted[i] = nums[i];
        Arrays.sort(sorted, (a, b) -> {
            int sa = digitSum(a), sb = digitSum(b);
            if (sa != sb) return sa - sb;
            return a - b;
        });
        Map<Integer, Integer> pos = new HashMap<>();
        for (int i = 0; i < n; i++) pos.put(sorted[i], i);
        boolean[] visited = new boolean[n];
        int cycles = 0;
        for (int i = 0; i < n; i++) {
            if (!visited[i]) {
                cycles++;
                int j = i;
                while (!visited[j]) {
                    visited[j] = true;
                    j = pos.get(nums[j]);
                }
            }
        }
        return n - cycles;
    }
    private int digitSum(int x) {
        int s = 0;
        while (x > 0) {
            s += x % 10;
            x /= 10;
        }
        return s;
    }
}