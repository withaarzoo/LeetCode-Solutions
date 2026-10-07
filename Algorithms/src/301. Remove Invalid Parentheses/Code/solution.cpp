class Solution {
public:
    vector<string> removeInvalidParentheses(string s) {
        int left = 0, right = 0;
        for (char c : s) {
            if (c == '(') left++;
            else if (c == ')') {
                if (left > 0) left--;
                else right++;
            }
        }
        unordered_set<string> res;
        string path;
        dfs(s, 0, left, right, 0, path, res);
        return vector<string>(res.begin(), res.end());
    }
private:
    void dfs(const string& s, int i, int left, int right, int open, string& path, unordered_set<string>& res) {
        if (i == s.size()) {
            if (left == 0 && right == 0 && open == 0) res.insert(path);
            return;
        }
        char c = s[i];
        if (c != '(' && c != ')') {
            path.push_back(c);
            dfs(s, i + 1, left, right, open, path, res);
            path.pop_back();
            return;
        }
        if (c == '(') {
            if (left > 0) {
                dfs(s, i + 1, left - 1, right, open, path, res);
            }
            path.push_back(c);
            dfs(s, i + 1, left, right, open + 1, path, res);
            path.pop_back();
        } else {
            if (right > 0) {
                dfs(s, i + 1, left, right - 1, open, path, res);
            }
            if (open > 0) {
                path.push_back(c);
                dfs(s, i + 1, left, right, open - 1, path, res);
                path.pop_back();
            }
        }
    }
};