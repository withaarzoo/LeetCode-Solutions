class Solution { 
    public List<String> generateParenthesis(int n) { 
        List<String> ans = new ArrayList<>();
        StringBuilder cur = new StringBuilder();
        dfs(n, n, cur, ans);
        return ans;
    }
    private void dfs(int open, int close, StringBuilder cur, List<String> ans) {
        if (open == 0 && close == 0) {
            ans.add(cur.toString());
            return;
        }
        if (open > 0) {
            cur.append('(');
            dfs(open - 1, close, cur, ans);
            cur.deleteCharAt(cur.length() - 1);
        }
        if (close > open) {
            cur.append(')');
            dfs(open, close - 1, cur, ans);
            cur.deleteCharAt(cur.length() - 1);
        }
    }
}