class Solution { 
public: 
    vector<string> generateParenthesis(int n) { 
        vector<string> ans;
        string cur;
        function<void(int,int)> dfs = [&](int open, int close) {
            if (open == 0 && close == 0) {
                ans.push_back(cur);
                return;
            }
            if (open > 0) {
                cur.push_back('(');
                dfs(open - 1, close);
                cur.pop_back();
            }
            if (close > open) {
                cur.push_back(')');
                dfs(open, close - 1);
                cur.pop_back();
            }
        };
        dfs(n, n);
        return ans;
    } 
};