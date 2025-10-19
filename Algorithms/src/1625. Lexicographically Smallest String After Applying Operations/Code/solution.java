import java.util.*;

class Solution {
    public String findLexSmallestString(String s, int a, int b) {
        int n = s.length();
        Set<String> seen = new HashSet<>();
        Queue<String> q = new ArrayDeque<>();
        seen.add(s);
        q.add(s);
        String ans = s;

        while (!q.isEmpty()) {
            String cur = q.poll();
            if (cur.compareTo(ans) < 0) ans = cur;

            // Operation 1: add a to odd indices (1,3,5,...)
            char[] arr = cur.toCharArray();
            for (int i = 1; i < n; i += 2) {
                int d = (arr[i] - '0' + a) % 10;
                arr[i] = (char)('0' + d);
            }
            String addOp = new String(arr);
            if (!seen.contains(addOp)) {
                seen.add(addOp);
                q.add(addOp);
            }

            // Operation 2: rotate right by b
            String rotOp = cur.substring(n - b) + cur.substring(0, n - b);
            if (!seen.contains(rotOp)) {
                seen.add(rotOp);
                q.add(rotOp);
            }
        }
        return ans;
    }
}
