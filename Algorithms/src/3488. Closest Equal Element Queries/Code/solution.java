class Solution {
    public List<Integer> solveQueries(int[] nums, int[] queries) {
        int n = nums.length;

        // Store all indices for every value
        Map<Integer, List<Integer>> positions = new HashMap<>();

        for (int i = 0; i < n; i++) {
            positions.computeIfAbsent(nums[i], k -> new ArrayList<>()).add(i);
        }

        int[] answer = new int[n];
        Arrays.fill(answer, -1);

        // Process each value group
        for (List<Integer> pos : positions.values()) {
            int m = pos.size();

            if (m == 1)
                continue;

            for (int i = 0; i < m; i++) {
                int curr = pos.get(i);

                int prev = pos.get((i - 1 + m) % m);
                int next = pos.get((i + 1) % m);

                int distPrev = Math.abs(curr - prev);
                distPrev = Math.min(distPrev, n - distPrev);

                int distNext = Math.abs(curr - next);
                distNext = Math.min(distNext, n - distNext);

                answer[curr] = Math.min(distPrev, distNext);
            }
        }

        List<Integer> result = new ArrayList<>();

        for (int idx : queries) {
            result.add(answer[idx]);
        }

        return result;
    }
}