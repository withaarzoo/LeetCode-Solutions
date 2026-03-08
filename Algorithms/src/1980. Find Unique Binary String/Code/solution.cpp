class Solution {
public:
    string findDifferentBinaryString(vector<string>& nums) {
        int n = nums.size();
        string result = "";

        // Construct a new string by flipping the diagonal bits
        for (int i = 0; i < n; i++) {
            // If current bit is '0', make it '1'
            // If current bit is '1', make it '0'
            if (nums[i][i] == '0')
                result += '1';
            else
                result += '0';
        }

        return result;
    }
};