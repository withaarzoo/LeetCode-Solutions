class Solution {
    public int findTheDistanceValue(int[] arr1, int[] arr2, int d) {
        Arrays.sort(arr2);
        int ans = 0;
        for (int x : arr1) {
            int idx = Arrays.binarySearch(arr2, x);
            if (idx < 0) idx = -idx - 1;
            boolean ok = true;
            if (idx < arr2.length && Math.abs(arr2[idx] - x) <= d) ok = false;
            if (idx > 0 && Math.abs(arr2[idx - 1] - x) <= d) ok = false;
            if (ok) ans++;
        }
        return ans;
    }
}