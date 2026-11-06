class Solution {
    public int[] countServers(int n, int[][] logs, int x, int[] queries) {
        Arrays.sort(logs, (a, b) -> Integer.compare(a[1], b[1]));
        int m = queries.length;
        int[][] q = new int[m][2];
        for (int i = 0; i < m; i++) {
            q[i][0] = queries[i];
            q[i][1] = i;
        }
        Arrays.sort(q, (a, b) -> Integer.compare(a[0], b[0]));
        int[] ans = new int[m];
        Map<Integer, Integer> freq = new HashMap<>();
        int active = 0;
        int left = 0, right = 0;
        int L = logs.length;
        for (int[] pair : q) {
            int time = pair[0];
            int idx = pair[1];
            while (right < L && logs[right][1] <= time) {
                int s = logs[right][0];
                freq.put(s, freq.getOrDefault(s, 0) + 1);
                if (freq.get(s) == 1) active++;
                right++;
            }
            while (left < right && logs[left][1] < time - x) {
                int s = logs[left][0];
                freq.put(s, freq.get(s) - 1);
                if (freq.get(s) == 0) active--;
                left++;
            }
            ans[idx] = n - active;
        }
        return ans;
    }
}