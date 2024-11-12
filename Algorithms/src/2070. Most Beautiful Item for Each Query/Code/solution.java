import java.util.*;

class Solution {
    public int[] maximumBeauty(int[][] items, int[] queries) {
        Arrays.sort(items, Comparator.comparingInt(a -> a[0]));

        List<int[]> priceBeauty = new ArrayList<>();
        int maxBeauty = 0;
        for (int[] item : items) {
            maxBeauty = Math.max(maxBeauty, item[1]);
            priceBeauty.add(new int[] { item[0], maxBeauty });
        }

        int[] result = new int[queries.length];
        for (int i = 0; i < queries.length; i++) {
            int query = queries[i];
            int idx = binarySearch(priceBeauty, query);
            result[i] = idx != -1 ? priceBeauty.get(idx)[1] : 0;
        }
        return result;
    }

    private int binarySearch(List<int[]> priceBeauty, int query) {
        int left = 0, right = priceBeauty.size() - 1;
        while (left <= right) {
            int mid = left + (right - left) / 2;
            if (priceBeauty.get(mid)[0] <= query)
                left = mid + 1;
            else
                right = mid - 1;
        }
        return right;
    }
}