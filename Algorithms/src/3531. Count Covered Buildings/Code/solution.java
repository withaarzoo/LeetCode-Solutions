import java.util.*;

class Solution {
    public int countCoveredBuildings(int n, int[][] buildings) {
        int m = buildings.length;
        Map<Integer, List<Integer>> row = new HashMap<>(); // x -> list of y
        Map<Integer, List<Integer>> col = new HashMap<>(); // y -> list of x

        // Build maps
        for (int[] b : buildings) {
            int x = b[0], y = b[1];
            row.computeIfAbsent(x, k -> new ArrayList<>()).add(y);
            col.computeIfAbsent(y, k -> new ArrayList<>()).add(x);
        }

        // Sort lists
        for (List<Integer> ys : row.values())
            Collections.sort(ys);
        for (List<Integer> xs : col.values())
            Collections.sort(xs);

        int ans = 0;
        for (int[] b : buildings) {
            int x = b[0], y = b[1];
            List<Integer> ys = row.get(x);
            List<Integer> xs = col.get(y);

            int posY = Collections.binarySearch(ys, y);
            int posX = Collections.binarySearch(xs, x);
            // binarySearch returns index >=0 for found
            boolean insideRow = (posY > 0 && posY < ys.size() - 1);
            boolean insideCol = (posX > 0 && posX < xs.size() - 1);

            if (insideRow && insideCol)
                ans++;
        }
        return ans;
    }
}
