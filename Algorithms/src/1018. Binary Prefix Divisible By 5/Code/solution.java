import java.util.*;

class Solution {
    public List<Boolean> prefixesDivBy5(int[] nums) {
        List<Boolean> ans = new ArrayList<>(nums.length);

        int rem = 0; // remainder of current prefix modulo 5

        for (int bit : nums) {
            // new value = old * 2 + bit (binary shift)
            rem = (rem * 2 + bit) % 5;

            // check divisibility by 5 using remainder
            ans.add(rem == 0);
        }

        return ans;
    }
}
