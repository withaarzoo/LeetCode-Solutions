class Solution {
    public String removeOuterParentheses(String s) {
        StringBuilder res = new StringBuilder();
        int bal = 0;
        for (char c : s.toCharArray()) {
            if (c == '(') {
                if (bal > 0) res.append(c);
                bal++;
            } else {
                bal--;
                if (bal > 0) res.append(c);
            }
        }
        return res.toString();
    }
}