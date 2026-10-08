class Solution {
public:
    string removeOuterParentheses(string s) {
        string res;
        int bal = 0;
        for (char c : s) {
            if (c == '(') {
                if (bal > 0) res += c;
                bal++;
            } else {
                bal--;
                if (bal > 0) res += c;
            }
        }
        return res;
    }
};