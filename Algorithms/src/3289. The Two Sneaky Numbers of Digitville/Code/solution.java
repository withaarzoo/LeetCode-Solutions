import java.util.*;

class Solution {
    public int[] getSneakyNumbers(int[] nums) {
        int n = nums.length - 2; // original range size
        boolean[] seen = new boolean[n]; // mark seen values
        int[] res = new int[2];
        int idx = 0;
        for (int x : nums) {
            if (seen[x]) {
                res[idx++] = x; // repeated -> add to result
                if (idx == 2)
                    break; // early stop once we have both
            } else {
                seen[x] = true; // mark seen
            }
        }
        return res;
    }
}
