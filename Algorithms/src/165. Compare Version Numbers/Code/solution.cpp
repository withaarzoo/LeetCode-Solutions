class Solution {
public:
    int compareVersion(string version1, string version2) {
        int i = 0, j = 0;
        int n = version1.size(), m = version2.size();

        while (i < n || j < m) {
            long num1 = 0, num2 = 0;
            // parse next number from version1
            while (i < n && version1[i] != '.') {
                num1 = num1 * 10 + (version1[i] - '0');
                ++i;
            }
            // skip dot if present
            if (i < n && version1[i] == '.') ++i;

            // parse next number from version2
            while (j < m && version2[j] != '.') {
                num2 = num2 * 10 + (version2[j] - '0');
                ++j;
            }
            if (j < m && version2[j] == '.') ++j;

            if (num1 < num2) return -1;
            if (num1 > num2) return 1;
        }
        return 0;
    }
};
