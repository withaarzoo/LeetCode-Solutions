import java.util.Arrays;

class Solution {
    public int[] successfulPairs(int[] spells, int[] potions, long success) {
        Arrays.sort(potions);               // sort potions
        int n = spells.length;
        int m = potions.length;
        int[] ans = new int[n];

        for (int i = 0; i < n; ++i) {
            long s = spells[i];                        // promote to long
            long need = (success + s - 1) / s;         // ceil(success / s)
            int pos = lowerBound(potions, need);       // first index >= need
            ans[i] = m - pos;
        }
        return ans;
    }

    // custom lowerBound because Arrays.binarySearch is less convenient for insertion pos
    private int lowerBound(int[] arr, long target) {
        int l = 0, r = arr.length;
        while (l < r) {
            int mid = (l + r) >>> 1;
            if ((long)arr[mid] < target) l = mid + 1;
            else r = mid;
        }
        return l;
    }
}
