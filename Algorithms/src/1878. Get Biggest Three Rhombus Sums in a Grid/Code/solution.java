class Solution {
    public int[] getBiggestThree(int[][] grid) {
        int m = grid.length, n = grid[0].length;
        TreeSet<Integer> set = new TreeSet<>();

        for (int r = 0; r < m; r++) {
            for (int c = 0; c < n; c++) {

                set.add(grid[r][c]);

                int maxSize = Math.min(Math.min(r, c), Math.min(m - 1 - r, n - 1 - c));

                for (int k = 1; k <= maxSize; k++) {
                    int sum = 0;

                    for (int i = 0; i < k; i++)
                        sum += grid[r - k + i][c + i];

                    for (int i = 0; i < k; i++)
                        sum += grid[r + i][c + k - i];

                    for (int i = 0; i < k; i++)
                        sum += grid[r + k - i][c - i];

                    for (int i = 0; i < k; i++)
                        sum += grid[r - i][c - k + i];

                    set.add(sum);
                }
            }
        }

        List<Integer> list = new ArrayList<>(set);
        Collections.sort(list, Collections.reverseOrder());

        int size = Math.min(3, list.size());
        int[] ans = new int[size];

        for (int i = 0; i < size; i++)
            ans[i] = list.get(i);

        return ans;
    }
}