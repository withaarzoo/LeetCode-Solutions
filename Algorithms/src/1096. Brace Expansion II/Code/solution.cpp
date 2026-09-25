class Solution {
public:
    vector<string> braceExpansionII(string expression) {
        int i = 0;
        set<string> res = parse(expression, i);
        return vector<string>(res.begin(), res.end());
    }
private:
    set<string> parse(const string& expr, int& i) {
        set<string> res;
        set<string> cur{""};
        while (i < expr.size() && expr[i] != '}') {
            if (expr[i] == '{') {
                ++i;
                set<string> next = parse(expr, i);
                ++i;
                cur = product(cur, next);
            } else if (expr[i] == ',') {
                res.insert(cur.begin(), cur.end());
                cur = {""};
                ++i;
            } else {
                set<string> next{string(1, expr[i])};
                ++i;
                cur = product(cur, next);
            }
        }
        res.insert(cur.begin(), cur.end());
        return res;
    }
    set<string> product(const set<string>& a, const set<string>& b) {
        set<string> res;
        for (const string& x : a)
            for (const string& y : b)
                res.insert(x + y);
        return res;
    }
};