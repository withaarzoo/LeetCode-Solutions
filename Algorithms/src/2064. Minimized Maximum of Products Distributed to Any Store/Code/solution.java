class Solution {
    private boolean canDistribute(int[] quantities, int maxProducts, int n) {
        int storesNeeded = 0;
        for (int quantity : quantities) {
            storesNeeded += (int) Math.ceil((double) quantity / maxProducts);
            if (storesNeeded > n)
                return false;
        }
        return storesNeeded <= n;
    }

    public int minimizedMaximum(int n, int[] quantities) {
        int low = 1;
        int high = Arrays.stream(quantities).max().getAsInt();
        int answer = high;

        while (low <= high) {
            int mid = low + (high - low) / 2;
            if (canDistribute(quantities, mid, n)) {
                answer = mid;
                high = mid - 1;
            } else {
                low = mid + 1;
            }
        }

        return answer;
    }
}