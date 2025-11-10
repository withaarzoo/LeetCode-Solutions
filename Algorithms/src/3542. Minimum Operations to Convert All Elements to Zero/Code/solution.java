import java.util.*;

class Solution {
    public int minOperations(int[] nums) {
        Deque<Integer> stk = new ArrayDeque<>(); // acts as a non-decreasing stack
        int ans = 0;
        for (int x : nums) {
            while (!stk.isEmpty() && stk.peekLast() > x)
                stk.pollLast();
            if (x == 0)
                continue;
            if (stk.isEmpty() || stk.peekLast() < x) {
                ans++;
                stk.addLast(x);
            }
        }
        return ans;
    }
}
