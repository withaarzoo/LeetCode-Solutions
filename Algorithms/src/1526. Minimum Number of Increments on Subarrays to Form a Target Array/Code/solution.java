class Solution {
    public int minNumberOperations(int[] target) {
        if (target == null || target.length == 0)
            return 0;
        int n = target.length;
        long ans = target[0]; // operations for the first element
        for (int i = 1; i < n; i++) {
            if (target[i] > target[i - 1]) {
                ans += (long) (target[i] - target[i - 1]); // add only positive increments
            }
        }
        return (int) ans; // problem guarantees result fits in 32-bit
    }
}
