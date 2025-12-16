import java.util.*;

class Solution {

    int B;
    List<Integer>[] tree;
    int[] present, future;
    int[][][] dp; // dp[node][parentBought][budget]

    public int maxProfit(int n, int[] present, int[] future,
            int[][] hierarchy, int budget) {

        this.B = budget;
        this.present = present;
        this.future = future;

        tree = new ArrayList[n];
        for (int i = 0; i < n; i++)
            tree[i] = new ArrayList<>();

        for (int[] e : hierarchy) {
            tree[e[0] - 1].add(e[1] - 1);
        }

        dp = new int[n][2][B + 1];

        dfs(0);

        int ans = 0;
        for (int b = 0; b <= B; b++) {
            ans = Math.max(ans, dp[0][0][b]);
        }
        return ans;
    }

    // Merge two knapsack arrays
    private int[] merge(int[] A, int[] B2) {
        int[] C = new int[B + 1];
        Arrays.fill(C, Integer.MIN_VALUE / 2);

        for (int i = 0; i <= B; i++) {
            if (A[i] < 0)
                continue;
            for (int j = 0; i + j <= B; j++) {
                C[i + j] = Math.max(C[i + j], A[i] + B2[j]);
            }
        }
        return C;
    }

    private void dfs(int u) {
        for (int v : tree[u]) {
            dfs(v);
        }

        for (int parentBought = 0; parentBought <= 1; parentBought++) {

            int price = parentBought == 1 ? present[u] / 2 : present[u];
            int profit = future[u] - price;

            // Option 1: skip buying u
            int[] skip = new int[B + 1];
            for (int v : tree[u]) {
                skip = merge(skip, dp[v][0]);
            }

            // Option 2: buy u
            int[] take = new int[B + 1];
            Arrays.fill(take, Integer.MIN_VALUE / 2);

            if (price <= B) {
                int[] base = new int[B + 1];
                for (int v : tree[u]) {
                    base = merge(base, dp[v][1]);
                }
                for (int b = price; b <= B; b++) {
                    take[b] = base[b - price] + profit;
                }
            }

            // Best of take / skip
            for (int b = 0; b <= B; b++) {
                dp[u][parentBought][b] = Math.max(skip[b], take[b]);
            }
        }
    }
}
