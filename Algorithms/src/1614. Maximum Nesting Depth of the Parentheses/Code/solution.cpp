class Solution {
public:
    int maxDepth(string s) {
        int depth = 0, maxDepth = 0;
        for (char c : s) {
            if (c == '(') {
                depth++;
                if (depth > maxDepth) maxDepth = depth;
            } else if (c == ')') {
                depth--;
            }
        }
        return maxDepth;
    }
};