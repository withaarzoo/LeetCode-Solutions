class Solution {
    public int maxOperations(String s) {
        long ans = 0L; // use long to avoid overflow in intermediate sum
        long ones = 0L;
        int n = s.length();
        for (int i = 0; i < n; ++i) {
            char c = s.charAt(i);
            if (c == '1') {
                ++ones;
            } else { // c == '0'
                if (i > 0 && s.charAt(i - 1) == '1')
                    ans += ones;
            }
        }
        return (int) ans;
    }
}
