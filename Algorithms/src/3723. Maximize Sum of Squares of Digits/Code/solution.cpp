class Solution {
public:
    string maxSumOfSquares(int num, int sum) {
        if (sum > 9LL * num) return "";
        int q = sum / 9;
        int r = sum % 9;
        string ans;
        ans.reserve(num);
        ans.append(q, '9');
        if (r) ans.push_back('0' + r);
        ans.append(num - q - (r > 0), '0');
        return ans;
    }
};