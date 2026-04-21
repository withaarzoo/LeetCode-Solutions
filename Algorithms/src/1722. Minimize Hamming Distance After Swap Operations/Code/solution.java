class Solution {
    int[] parent;
    int[] rank;

    private int find(int x) {
        if (parent[x] != x) {
            parent[x] = find(parent[x]);
        }
        return parent[x];
    }

    private void union(int a, int b) {
        int pa = find(a);
        int pb = find(b);

        if (pa == pb)
            return;

        if (rank[pa] < rank[pb]) {
            parent[pa] = pb;
        } else if (rank[pb] < rank[pa]) {
            parent[pb] = pa;
        } else {
            parent[pb] = pa;
            rank[pa]++;
        }
    }

    public int minimumHammingDistance(int[] source, int[] target, int[][] allowedSwaps) {
        int n = source.length;

        parent = new int[n];
        rank = new int[n];

        for (int i = 0; i < n; i++) {
            parent[i] = i;
        }

        // Build connected components
        for (int[] swap : allowedSwaps) {
            union(swap[0], swap[1]);
        }

        // Group indices by root parent
        Map<Integer, List<Integer>> groups = new HashMap<>();

        for (int i = 0; i < n; i++) {
            int root = find(i);
            groups.putIfAbsent(root, new ArrayList<>());
            groups.get(root).add(i);
        }

        int answer = 0;

        // Process each component
        for (List<Integer> indices : groups.values()) {
            Map<Integer, Integer> freq = new HashMap<>();

            // Count source values
            for (int idx : indices) {
                freq.put(source[idx], freq.getOrDefault(source[idx], 0) + 1);
            }

            // Match target values
            for (int idx : indices) {
                if (freq.getOrDefault(target[idx], 0) > 0) {
                    freq.put(target[idx], freq.get(target[idx]) - 1);
                } else {
                    answer++;
                }
            }
        }

        return answer;
    }
}