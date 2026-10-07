class Solution {
    public List<String> removeInvalidParentheses(String s) {
        int left = 0, right = 0;
        for (char c : s.toCharArray()) {
            if (c == '(') left++;
            else if (c == ')') {
                if (left > 0) left--;
                else right++;
            }
        }
        Set<String> res = new HashSet<>();
        StringBuilder path = new StringBuilder();
        dfs(s, 0, left, right, 0, path, res);
        return new ArrayList<>(res);
    }
    private void dfs(String s, int i, int left, int right, int open, StringBuilder path, Set<String> res) {
        if (i == s.length()) {
            if (left == 0 && right == 0 && open == 0) res.add(path.toString());
            return;
        }
        char c = s.charAt(i);
        if (c != '(' && c != ')') {
            path.append(c);
            dfs(s, i + 1, left, right, open, path, res);
            path.deleteCharAt(path.length() - 1);
            return;
        }
        if (c == '(') {
            if (left > 0) {
                dfs(s, i + 1, left - 1, right, open, path, res);
            }
            path.append(c);
            dfs(s, i + 1, left, right, open + 1, path, res);
            path.deleteCharAt(path.length() - 1);
        } else {
            if (right > 0) {
                dfs(s, i + 1, left, right - 1, open, path, res);
            }
            if (open > 0) {
                path.append(c);
                dfs(s, i + 1, left, right, open - 1, path, res);
                path.deleteCharAt(path.length() - 1);
            }
        }
    }
}